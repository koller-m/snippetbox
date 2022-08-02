package main

import (
	"log"
	"net/http"
)

func main() {
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

	// Use http.ListenAndServe() to start a web server
	// Pass two params
	// TCP network address (":4000")
	// And the servemux we just created
	// If http.ListenAndServe() returns an error
	// We use log.Fatal() to log the error and exit
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
