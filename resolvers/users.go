package resolvers

import (
	"basic-api/models"
	"time"
)

func (u models.User) CreateUser() {
	db.Create(&models.User{
		FirstName: "Theo",
		LastName:  "Fenique",
		Dob:       time.Now().Format("2006-01-02"),
		City:      "Paris",
		Username:  "Kazurdan",
		Password:  "pass",
	})
}
