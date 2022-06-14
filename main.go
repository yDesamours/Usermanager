package main

import (
	"net/http"
	"time"
	"usermanager/dao"
	"usermanager/routes"
	"usermanager/sessionHandlers"
)

func main() {
	//try connecting to the database
	dao.Connect()
	sessionHandlers.Sessionstore.Cleanup(time.Hour)

	http.Handle("/", routes.Router())
	//start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Something Went wrong")
	}
}
