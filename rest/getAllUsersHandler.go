package rest

import (
	"encoding/json"
	"net/http"
	"usermanager/services"
)

//to get all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//query the dao for all users
	users, _ := services.GetAllUsers()
	//send response as json
	json.NewEncoder(w).Encode(users)

}
