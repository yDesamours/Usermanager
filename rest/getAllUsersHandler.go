package rest

import (
	"encoding/json"
	"net/http"
	"usermanager/services"
)

// GetUsers godoc
// @Description Get all the users
// @Accept  json
// @Success 200 array models.User "List of all the users"
// @Failure 401 {string} string
// @Router /api/usermanager/users [get]
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//query the dao for all users
	users, _ := services.GetAllUsers()
	//send response as json
	json.NewEncoder(w).Encode(users)

}
