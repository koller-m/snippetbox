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

	// Init a slice containing the paths to two files
	// The file containing the base template must be first
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Use template.ParseFiles() to read the files and store the templates
	// Log error with http.Error() to send a generic 500 Internal Server Error
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Use ExecuteTemplate() to write the content of the "base" template
	// As the response body
	err = ts.ExecuteTemplate(w, "base", nil)
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
