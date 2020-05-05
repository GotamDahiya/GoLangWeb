package main

import (
	"fmt"
	"net/http"
	"log"

	dbcon "./dbcon"
)

// ServeStatic-> serves static files and gets input from those files
// func ServeStatic(w http.ResponseWriter,r *http.Request) {
// 	http.fileServer(http.Dir("./static"))
// 	err := r.ParseForm()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	name := r.PostFormValue("name")
// 	fmt.Println( name)
// }

func main() {
	dbcon.ConnectDB()

	http.Handle("/", http.FileServer(http.Dir("./static"))) // http.HandleFunc("/", ServeStatic)

	http.HandleFunc("/hi",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	
	log.Fatal(http.ListenAndServe(":8000", nil))
}