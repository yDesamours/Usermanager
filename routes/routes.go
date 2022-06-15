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
	router.HandleFunc("/api/usermanager/users", sessionHandlers.IsloggedInHandler(sessionHandlers.HaveRight(rest.GetUsersHandler, "view users"))).Methods("GET")
	router.HandleFunc("/api/usermanager/updateuser", sessionHandlers.IsloggedInHandler(sessionHandlers.HaveRight(rest.EditUserHandler, "edit self"))).Methods("PUT")
	router.HandleFunc("/api/usermanager/updatepassword", sessionHandlers.HaveRight(rest.EditPasswordHandler, "edit password")).Methods("PUT")
	router.HandleFunc("/api/usermanager/adminupdateuser", sessionHandlers.HaveRight(rest.AdminEditUserHandler, "edit user")).Methods("PUT")
	router.HandleFunc("/api/usermanager/login", sessionHandlers.LoginHandler).Methods("POST")
	router.HandleFunc("/api/usermanager/logout", sessionHandlers.IsloggedInHandler(sessionHandlers.LogoutHandler)).Methods("GET")

	return router
}
