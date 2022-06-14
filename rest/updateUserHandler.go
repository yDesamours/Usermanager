package rest

import (
	"fmt"
	"net/http"
	"usermanager/services"
	"usermanager/sessionHandlers"
)

//allow a user to change his password
func EditPasswordHandler(w http.ResponseWriter, r *http.Request) {

	//get the current user info
	currentUser, _ := sessionHandlers.GetUser(r)

	//query the dao to change the password
	if update := services.EditpasswordService(r.Body, *currentUser); update == nil {
		fmt.Fprintf(w, "Password Updated!")
	} else {
		fmt.Fprintf(w, update.Error())
	}
}
