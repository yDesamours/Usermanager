package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/database"
	"usermanager/models"
	"usermanager/sessionHandlers"
	"usermanager/utils"
)

//for adding a new user
func AddUserHandler(w http.ResponseWriter, r *http.Request) {

	//etract the json data from the request. copy them into a user structure
	var newUser models.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	//get the information about the current user
	currentUser, err := sessionHandlers.GetUser(r)

	//set de user's role and activity to their default value
	if err != nil || currentUser.Role != "admin" {
		newUser.Role = "client"
		newUser.IsActive = true
	}

	//test user credentials
	if err := utils.TestCredentials(newUser, true); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	//hash the submitted password
	newUser.Password = utils.HashPassword(newUser.Password)
	//lowercase all character, expect the password
	utils.Sanitize(&newUser)
	//insert thw user
	result := database.InsertUser(newUser)
	fmt.Println(newUser)
	//the insertion may failed
	//the result old info describing what happens
	fmt.Fprintf(w, result)
}

//to get all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//query the database for all users
	users := database.GetAllUsers()
	//send response as json
	json.NewEncoder(w).Encode(users)

}

//Allow a user to modify his own data
func EditUserHandler(w http.ResponseWriter, r *http.Request) {

	//extract new info from request body
	var edit models.User
	json.NewDecoder(r.Body).Decode(&edit)

	//get the current user infos
	currentUser, _ := sessionHandlers.GetUser(r)

	//thes for password matching. On failure, end the process
	if ok := utils.ComparePassword(edit.Password, currentUser.Password); !ok {
		fmt.Fprintf(w, "Incorect password")
		return
	}

	//test for credentials. On failure, end the proces
	if err := utils.TestCredentials(edit, false); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//lowercase everything
	utils.Sanitize(&edit)
	//query the database for update
	if update := database.EditUser(currentUser.Username, edit); update {
		fmt.Fprintf(w, "Username updated to %s", edit.Username)
	} else {
		fmt.Fprintf(w, "Failed to update username")
	}
}

//allow a user to change his password
func EditPasswordHandler(w http.ResponseWriter, r *http.Request) {

	//get old password and new password submitted by the user
	edit := struct {
		ActualPassword string `json:"password"`
		NewPassword    string `json:"newpassword"`
	}{}
	json.NewDecoder(r.Body).Decode(&edit)
	//get the current user info
	currentUser, _ := sessionHandlers.GetUser(r)
	//test for password matching
	if ok := utils.ComparePassword(edit.ActualPassword, currentUser.Password); !ok {
		fmt.Fprintf(w, "Incorect password")
		return
	}
	//test the new password
	if result, ok := utils.TestPassword(edit.NewPassword); !ok {
		fmt.Fprintf(w, result)
		return
	}
	//hash the password
	hash := utils.HashPassword(edit.NewPassword)
	//query the database to change the password
	if update := database.EditPassword(hash, currentUser.Username); update {
		fmt.Fprintf(w, "Password Updated!")
	} else {
		fmt.Fprintf(w, "Failed to update password")
	}
}

//this handler permits an admin to modify a user's info
func AdminEditUserHandler(w http.ResponseWriter, r *http.Request) {
	//get the current user
	currentUser, _ := sessionHandlers.GetUser(r)
	//if he is not an admin, stop the process
	if currentUser.Role != "admin" {
		fmt.Fprintf(w, "Not allowed")
		return
	}

	//etract the json data from the request. Copy them into a structure
	//This structure has 2 fields. One for storing new data about the user, but the password sent his the admin's
	//password. The other fields hold the username of the user whose data must be updated
	editor := struct {
		User     models.User `json:"user"`
		Username string      `json:"targetusername"`
	}{}
	json.NewDecoder(r.Body).Decode(&editor) //all data are strored.fmt.Println(editor)
	//call the database to retrieve informations about the targeted user
	targetedUser, err := database.GetUser(editor.Username)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//If the targeted user is an admin, stop the process
	if targetedUser.Role == "admin" {
		fmt.Fprintf(w, "Not allowed")
		return
	}

	//check for credentials
	//If the test failed, stop the process and return a message describing what's wrong
	if err := utils.TestCredentials(editor.User, false); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//test the password of the user
	//if the test failed, what do we do? Stop the process. On connait la chanson
	if ok := utils.ComparePassword(editor.User.Password, currentUser.Password); !ok {
		fmt.Fprintf(w, "Incorect password")
		return
	}
	//everything is correct. Contact the database to make the change
	if ok := database.AdminEditUser(editor.Username, editor.User); !ok {
		fmt.Fprintf(w, "Failed to update user's infos!")
	} else {
		fmt.Fprintf(w, "User's infos updated!")
	}

}
