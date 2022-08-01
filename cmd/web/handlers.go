package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

	// Use template.ParseFiles() to read template file into template set
	// Log error with http.Error() to send a generic 500 Internal Server Error
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Use Execute() method on template set to write template content as
	// Response body
	// The last param of Execute() represents any dynamic data we want to pass in
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Add snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id param from query string
	// Convert it to int using strconv.Atoi()
	// If it can't be converted or value is less than 1, return 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use fmt.Fprintf() to interpolate the id value with our response
	// and write it to http.ResponseWriter
	fmt.Fprintf(w, "Display a specific snippet with the ID %d...", id)
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
