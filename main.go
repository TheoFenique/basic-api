package main

import (
	"basic-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", controllers.GetArticles).Methods("GET")
	r.HandleFunc("/articles", controllers.PostArticles).Methods("POST")
}
