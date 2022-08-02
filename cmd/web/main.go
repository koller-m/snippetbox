package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a command-line flag "addr"
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Parse the command-line flag with flag.Parse()
	// This reads in the command-line flag and assigns it to addr
	// Must be called before the addr variable is used
	flag.Parse()

	// Use http.NewServeMux() to init new servemux
	// Register the home func as the handler for the "/" URL pattern
	mux := http.NewServeMux()

	// Create file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use mux.Handle() to register the file server as the handler for
	// All URL paths that start with "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// The value returned from flag.String() is a pointer to the flag value
	// NOT the value itself
	// Dereference the pointer with the * symbol
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
