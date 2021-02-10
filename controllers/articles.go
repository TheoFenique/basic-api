package controllers

import (
	"encoding/json"
	"net/http"
)

// GetArticles enocodes all the articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	article := struct {
		ID   string
		text string
	}{
		"333",
		"Texttttt",
	}

	json.NewEncoder(w).Encode(article)
}

// PostArticles posts an array of articles into the database
func PostArticles(w http.ResponseWriter, r *http.Request) {
	// fill database
}
