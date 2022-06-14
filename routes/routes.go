package routes

import (
	"usermanager/rest"
	"usermanager/sessionHandlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	//create a router
	router := mux.NewRouter()
	//mount all the routes
	router.HandleFunc("/api/usermanager/register", rest.AddUserHandler).Methods("POST")
	router.HandleFunc("/api/usermanager/users", sessionHandlers.IsloggedInHandler(rest.GetUsersHandler)).Methods("GET")
	router.HandleFunc("/api/usermanager/updateuser", sessionHandlers.IsloggedInHandler(rest.EditUserHandler)).Methods("PUT")
	router.HandleFunc("/api/usermanager/updatepassword", sessionHandlers.IsloggedInHandler(rest.EditPasswordHandler)).Methods("PUT")
	router.HandleFunc("/api/usermanager/adminupdateuser", sessionHandlers.IsloggedInHandler(rest.AdminEditUserHandler)).Methods("PUT")
	router.HandleFunc("/api/usermanager/login", sessionHandlers.LoginHandler).Methods("POST")
	router.HandleFunc("/api/usermanager/logout", sessionHandlers.IsloggedInHandler(sessionHandlers.LogoutHandler)).Methods("GET")

	return router
}
