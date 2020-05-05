package dbcon

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //used for connecting to the MySQL database
)

var db *sql.DB
var err error

// ConnectDB -> connecting to the MySQL database
func ConnectDB() {
	fmt.Println("Database Connected")
	db, err = sql.Open("mysql","root:Abcd@1234@/server_go")
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}