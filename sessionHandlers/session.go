package sessionHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/dao"
	"usermanager/models"
	"usermanager/utils"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
)

var connectionString = dao.CreateConnectionString()
var Sessionstore, err = pgstore.NewPGStore(connectionString, []byte("my secure key"))

// Adduser godoc
// @Description Allows a registered user to login
// @Accept  json
// @Success 200 {object} string "Login succeed"
// @Failure 403 {string} string "Login failed"
// @Router /api/usermanager/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var logger models.User //to store user credentials
	//copy the user credentials sent as json
	json.NewDecoder(r.Body).Decode(&logger)
	//search for the user in the dao
	savedUser, err := dao.GetUser(logger.Username)
	//if user does not exist

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, err.Error())
		return
	}

	//if password is incorrect
	if ok := utils.ComparePassword(logger.Password, savedUser.Password); !ok {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Incorrect password!")
		return
	}
	//create or retrieve a session
	session, _ := Sessionstore.Get(r, "current-session")
	//set session maxAge
	session.Options = &sessions.Options{
		MaxAge: 3600,
	}
	//store the username in the session.Values interface
	session.Values["username"] = logger.Username
	//save the session
	session.Save(r, w)
	//respond the user
	fmt.Fprintf(w, "You are logged in %s!", session.Values["username"])

}

// Adduser godoc
// @Description End a session
// @Success 200 {object} string "Logout succeed"
// @Router /api/usermanager/logout [post]
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//get the session
	session, _ := Sessionstore.Get(r, "currentsession")
	//delete teh username key and save the session
	delete(session.Values, "username")
	session.Save(r, w)
	//respond the user
	fmt.Fprintf(w, "You are logged out!")
	//the user is logged out
}

//this function retrieve the infos about the connected user
func GetUser(r *http.Request) (*models.User, error) {

	//get the session
	session, _ := Sessionstore.Get(r, "current-session")
	//the username key stored in the session.Values interface is an interface{}
	//convert in into a string
	username := fmt.Sprint(session.Values["username"])
	//query the dao to get the user info
	currentUser, err := dao.GetUser(username)

	return currentUser, err
}

//Middleware that check if a user is log in before accessing a route
func IsloggedInHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Sessionstore.Get(r, "current-session")
		//if the user is not logged in, end the process
		if _, ok := session.Values["username"]; !ok {
			http.Redirect(w, r, "api/usermanager/login", http.StatusUnauthorized)
			return
		}
		//if user is connected, let him access the desired route
		f(w, r)

	}
}

//middleware that check if the have the correct right to access this route.
//Required permission is passed as the second parameter
func HaveRight(f http.HandlerFunc, permission string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Sessionstore.Get(r, "current-session")
		//if the user is not logged in, end the process
		if _, ok := session.Values["username"]; !ok {
			http.Redirect(w, r, "api/usermanager/login", http.StatusUnauthorized)
			return
		}
		//only admin can access this route
		var allowed bool
		//get the currentuser infos
		currentUser, _ := GetUser(r)
		//get all the user's permissions
		permissions, _ := dao.GetRolePermissions(currentUser.Role)
		for _, permission := range permissions {
			fmt.Println(permission)
			if permission == permission {
				allowed = true
				break
			}
		}
		//does the user have correct permisions
		if !allowed {
			fmt.Fprintf(w, "--Not allowed!")
			return
		}
		f(w, r)
	}
}
