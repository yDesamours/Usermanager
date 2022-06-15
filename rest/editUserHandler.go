package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/models"
	"usermanager/services"
	"usermanager/sessionHandlers"
)

//Allow a user to modify his own data
func EditUserHandler(w http.ResponseWriter, r *http.Request) {

	//extract new info from request body
	var edit models.User
	json.NewDecoder(r.Body).Decode(&edit)

	//get the current user infos
	currentUser, _ := sessionHandlers.GetUser(r)

	if update := services.EditUserService(&edit, currentUser); update == nil {
		fmt.Fprintf(w, "User's informations successfully updated!")
	} else {
		fmt.Fprintf(w, update.Error())
	}
}
