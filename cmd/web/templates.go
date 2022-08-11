package main

import (
	"html/template"
	"path/filepath"

	"github.com/koller-m/snippetbox/internal/models"
)

// Define templateData type to hold dynamic data for HTML templates
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Init new map to act as the cache
	cache := map[string]*template.Template{}

	// Use filepath.Glob() to get a slice of all filepaths that match
	// The pattern "./ui/html/pages/*.tmpl.html"
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths
	for _, page := range pages {
		// Extract the file name and assign it to name variable
		name := filepath.Base(page)

		// Parse the base template file into a template set
		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() on any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() on page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map
		// Name of page as key ("home.tmpl.html")
		cache[name] = ts
	}
	// Return the map
	return cache, nil
}
