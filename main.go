package main

import (
	"net/http"
	"usermanager/database"
	"usermanager/routes"
)

func main() {
	//try connecting to the database
	database.Connect()

	http.Handle("/", routes.Router())
	//start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Something Went wrong")
	}
}
