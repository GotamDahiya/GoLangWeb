package dbcon



//AddUser -> Adds the user's details to the MySQL databse
func AddUser(username, pwd, name, email, interests, hobbies, achivements string) {
	// fmt.Println(username+" "+pwd+" "+name+" "+email+" "+interests+" "+hobbies+" "+achivements)
	stmt, err := db.Prepare("insert into User_Profile(UserName,Password,Name,Email,Hobbies,Achivements,Interests) values(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(username,pwd,name,email,hobbies,achivements,interests)
	if err != nil {
		panic(err.Error())
	}
}