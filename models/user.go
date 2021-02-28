package models

// DBUser represents the user in the database
type DBUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Dob       string `json:"dob"`
	City      string `json:"city"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// LoggedUser is used to return user info
type LoggedUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Dob       string `json:"dob"`
	City      string `json:"city"`
	Username  string `json:"username"`
}
