package sessionHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/database"
	"usermanager/models"
	"usermanager/utils"

	"github.com/gorilla/sessions"
)

var Sessionstore = sessions.NewCookieStore([]byte("My very secret key"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var logger models.User //to store user credentials
	//copy the user credentials sent as json
	json.NewDecoder(r.Body).Decode(&logger)
	fmt.Println(logger)
	//search for the user in the database
	savedUser, message := database.GetUser(logger.Username)
	//if user does not exist

	if message != "" {
		fmt.Fprintf(w, "Username does not exist!")
		return
	}

	//if password is incorrect
	if ok := utils.ComparePassword(logger.Password, savedUser.Password); !ok {
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
func GetUser(r *http.Request) models.User {

	//get the session
	session, _ := Sessionstore.Get(r, "current-session")
	//the username key stored in the session.Values interface is an interface{}
	//convert in into a string
	username := fmt.Sprint(session.Values["username"])
	//query the database to get the user info
	currentUser, _ := database.GetUser(username)

	return currentUser
}

//Middlewarre to check if user is connected
func IsloggedInHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Sessionstore.Get(r, "current-session")
		//if the user is not logged in, end the process
		if _, ok := session.Values["username"]; !ok {
			http.Redirect(w, r, "api/usermanager/login", http.StatusForbidden)
			return
		}
		//if user is connected, let him access the desired route
		f(w, r)

	}
}
