package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Create application struct to hold application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a command-line flag "addr"
	addr := flag.String("addr", ":4000", "HTTP network address")

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

	// Init new instance of our application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Use http.NewServeMux() to init new servemux
	// Register the home func as the handler for the "/" URL pattern
	mux := http.NewServeMux()

	// Create file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use mux.Handle() to register the file server as the handler for
	// All URL paths that start with "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// Init a new http.Server struct
	// ErrorLog now uses the custom errorLog logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// Call ListenAndServe() on new http.Server struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
