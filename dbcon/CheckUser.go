package dbcon

import (
	"fmt"
)

// UserDetails -> obtain username and password
type UserDetails struct {
	uname string
	pwd string
}

var check int = 0

// CheckUser -> Checks if the user exists in the database or not
func CheckUser(uname, pwd string) int {
	fmt.Println(uname+" "+pwd)
	var (
		user UserDetails
		users []UserDetails
	)
	
	rows, err := db.Query("select UserName, Password from User_Profile where UserName = "+uname)
	if err != nil {
		panic(err.Error())
	}
	
	for rows.Next() {
		rows.Scan(&user.uname,&user.pwd)
		users = append(users,user)	
	}

	if len(users) >= 1 {
		check = 2
	}
	if pwd == users[0].pwd {
		check = 1
	} else {
		check = 0
	}
	return check
}