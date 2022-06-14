package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/models"
	"usermanager/services"
	"usermanager/sessionHandlers"
)

//for adding a new user
func AddUserHandler(w http.ResponseWriter, r *http.Request) {

	//etract the json data from the request. copy them into a user structure
	var newUser models.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	//get the information about the current user
	currentUser, err := sessionHandlers.GetUser(r)

	//set de user's role and activity to their default value
	if err != nil || currentUser.Role != "admin" {
		newUser.Role = "client"
		newUser.IsActive = false
	}
	//contact the services for insertion
	result := services.AddUser(newUser)
	//the insertion may fail
	if result != nil {
		fmt.Fprintf(w, result.Error())
	}
	//everything is ok
	fmt.Fprintf(w, "New user Inserted")
}
