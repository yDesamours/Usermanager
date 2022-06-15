package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/models"
	"usermanager/services"
	"usermanager/sessionHandlers"
)

// Adduser godoc
// @Description Create a user
// @Accept  json
// @param user body models.User true "The firstname of the person"
// @Success 200 {object} string "Insert a new user"
// @Failure 403 {string} string
// @Router /api/usermanager/register [post]
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
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, result.Error())
	}
	//everything is ok
	fmt.Fprintf(w, "New user Inserted")
}
