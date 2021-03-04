package controllers

import (
	"basic-api/config"
	"basic-api/models"
	"encoding/json"
	"net/http"
	"time"
)

// HandleLogin checks the email and password and returns the token
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// do stuff
}

// HandleRegister takes the UserForm, registers the user
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	db.Create(&models.User{
		FirstName: "Theo",
		LastName:  "Fenique",
		Dob:       time.Now().Format("2006-01-02"),
		City:      "Paris",
		Username:  "Kazurdan",
		Password:  "pass",
	})

	json.NewEncoder(w).Encode("done")
}

// GetProfile takes the UserForm, edits the profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// do stuff
}

// EditProfile takes the UserForm, edits the profile
func EditProfile(w http.ResponseWriter, r *http.Request) {
	// do stuff
}
