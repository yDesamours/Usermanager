package rest

import (
	"fmt"
	"net/http"
	"usermanager/services"
	"usermanager/sessionHandlers"
)

//this handler permits an admin to modify a user's info
func AdminEditUserHandler(w http.ResponseWriter, r *http.Request) {
	//get the current user
	currentUser, _ := sessionHandlers.GetUser(r)
	//if he is not an admin, stop the process
	if currentUser.Role != "admin" {
		fmt.Fprintf(w, "Not allowed")
		return
	}

	//everything is correct. Contact the dao to make the change
	if err := services.AdminUpdateUser(r.Body, *currentUser); err != nil {
		fmt.Fprintf(w, "Failed to update user's infos!")
	} else {
		fmt.Fprintf(w, "User's infos updated!")
	}

}
