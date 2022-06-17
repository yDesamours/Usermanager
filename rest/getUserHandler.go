package rest

import (
	"encoding/json"
	"net/http"
	"strings"
	"usermanager/services"
)

// GetUsers godoc
// @Description Get all the users
// @Accept  json
// @Success 200 array models.UserResponse "List of all the users"
// @Failure 401 {string} string
// @Router /api/usermanager/users [get]
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	username := path[strings.LastIndex(path, "/"):]
	//query the dao for all users
	users, _ := services.GetUser(username)
	//send response as json
	json.NewEncoder(w).Encode(users)

}
