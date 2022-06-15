package rest

import (
	"fmt"
	"net/http"
	"usermanager/services"
	"usermanager/sessionHandlers"
)

// EditPassword godoc
// @Description Create a user
// @Accept  json
// @Success 200 {object} string "Update a password"
// @Failure 401 {string} string
// @Router /api/usermanager/updatepassword [put]
func EditPasswordHandler(w http.ResponseWriter, r *http.Request) {

	//get the current user info
	currentUser, _ := sessionHandlers.GetUser(r)

	//query the dao to change the password
	if update := services.EditpasswordService(r.Body, *currentUser); update == nil {
		fmt.Fprintf(w, "Password Updated!")
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, update.Error())
	}
}