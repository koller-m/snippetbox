package main

import "github.com/koller-m/snippetbox/internal/models"

// Define templateData type to hold dynamic data for HTML templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
