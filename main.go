package main

import (
	"basic-api/config"
	"basic-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/article", controllers.PostArticles).Methods("POST")
	r.HandleFunc("/article/{id}", controllers.GetArticle).Methods("GET")
	r.HandleFunc("/article/{id}/like", controllers.LikeArticle).Methods("POST")
	r.HandleFunc("/article/{id}/dislike", controllers.DislikeArticle).Methods("POST")

	r.HandleFunc("/article/{id}/comment", controllers.PostComment).Methods("POST")
	r.HandleFunc("/article/{id}/comment/like", controllers.PostLikeComment).Methods("POST")
	r.HandleFunc("/article/{id}/comment/dislike", controllers.PostDislikeComment).Methods("POST")

	r.HandleFunc("/articles", controllers.GetArticles).Methods("GET")

	r.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	r.HandleFunc("/register", controllers.HandleRegister).Methods("POST")
	r.HandleFunc("/profile", controllers.GetProfile).Methods("GET")
	r.HandleFunc("/profile", controllers.EditProfile).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
