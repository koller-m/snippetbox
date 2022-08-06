package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Create application struct to hold application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a command-line flag "addr"
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Define command-line flag for MySQL DSN string
	// REMOVE pass
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	// Parse the command-line flag with flag.Parse()
	// This reads in the command-line flag and assigns it to addr
	// Must be called before the addr variable is used
	flag.Parse()

	// Create a logger for writing information messages. This takes three params
	// The destination to write the logs to, a string prefix for message
	// And flags to indicate what additional information to include
	// Flags are joined using the | operator
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Close the connection pool before main() exits
	defer db.Close()

	// Init new instance of our application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Init a new http.Server struct
	// ErrorLog now uses the custom errorLog logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB() wraps sql.Open() and returns a sql.DB connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
