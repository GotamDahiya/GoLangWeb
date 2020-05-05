package common

import (
	"fmt"
	"net/http"

	servers "../servers"
	dbcon "../dbcon"
)

//User -> to get detail of user and store them in the MySQL database
type User struct {
	Username string
	Password string
	Name string
	Hobbies string
	Email string
	Achivements string
	Interests string
}

// RegisterPageHandler -> Loads the register page
func RegisterPageHandler(w http.ResponseWriter,r *http.Request) {
	var body,_ = servers.LoadFile("static/register.html")
	fmt.Fprintf(w, body)
}

// RegisterHandler -> gets the values from the user through the register.html page
func RegisterHandler(w http.ResponseWriter,r *http.Request) {
	var (
		user User
		users []User
	)
	
}