package dbcon

import (
	"database/sql"
)


// ExecuteQuery -> executes search query entered by user
func ExecuteQuery(area string,query string) (*sql.Rows) {
	// fmt.Println("Hello World")
	rows,err := db.Query("select UserName, Name, Email, Hobbies, Achivements, Interests from User_Profile where "+area+" like '%"+query+"%'")
	if err != nil {
		panic(err)
	}
	return rows
}