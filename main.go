package main

import (
	"net/http"
	"time"
	"usermanager/dao"
	_ "usermanager/docs"
	"usermanager/routes"
	"usermanager/sessionHandlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title USERMANAGER API
// @version 1.0
// @description This is a mini api that allows to perform basics CRUD operations
// @contactYvelt DESAMOURS
// @contact.email dyvelt@tainosystems.com
// @host 192.168.10.137:9090
// @BasePath /api/usermanager
func main() {
	//try connecting to the database
	dao.Connect()
	sessionHandlers.Sessionstore.Cleanup(time.Hour)
	routes.Router().PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	http.Handle("/", routes.Router())

	//start the server
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic("Something Went wrong")
	}
}
