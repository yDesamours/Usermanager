package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/models"
	"usermanager/services"
	"usermanager/sessionHandlers"
)

// EditUser godoc
// @Description Edit a user's informations
// @Accept  json
// @Success 200 {object} string "Infos are edited"
// @Failure 403 {string} string
// @Failure 401 {string} string
// @Router /api/usermanager/updateuser [put]
func EditUserHandler(w http.ResponseWriter, r *http.Request) {

	//extract new info from request body
	var edit models.User
	json.NewDecoder(r.Body).Decode(&edit)

	//get the current user infos
	currentUser, _ := sessionHandlers.GetUser(r)

	if update := services.EditUserService(&edit, currentUser); update == nil {
		fmt.Fprintf(w, "User's informations successfully updated!")
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, update.Error())
	}
}
