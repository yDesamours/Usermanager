package rest

import (
	"fmt"
	"net/http"
	"usermanager/services"
	"usermanager/sessionHandlers"
)

// AdminEditUser godoc
// @Description Allows an admin yo edit a user's information
// @Accept  json
// @Success 200 {object} string "The user's info are edited"
// @Failure 403 {string} string
// @Failure 401 {string} string
// @Router /api/usermanager/adminupdateuser [put]
func AdminEditUserHandler(w http.ResponseWriter, r *http.Request) {
	//get the current user
	currentUser, _ := sessionHandlers.GetUser(r)
	//if he is not an admin, stop the process

	//everything is correct. Contact the dao to make the change
	if err := services.AdminUpdateUser(r.Body, *currentUser); err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "User's infos updated!")
	}

}
