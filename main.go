package main

import (
	"log"
	"net/http"
)

// Define home handler function
// Writes a byte slice containing
// "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request matches "/"
	// If not, use http.NotFound() to send 404 response
	// Return from the handler
	// If not, the handler would keep executing
	// And also write "Hello from Snippetbox"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check if the request is POST
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		// Use http.Error() to send 405 status code and "Method Not Allowed"
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use http.NewServeMux() to init new servemux
	// Register the home func as the handler for the "/" URL pattern
	mux := http.NewServeMux()
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
