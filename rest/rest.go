package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/dao"
	"usermanager/models"
	"usermanager/services"
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
		newUser.IsActive = false
	}
	//contact the services for insertion
	result := services.AddUser(newUser)
	//the insertion may fail
	if result != nil {
		fmt.Fprintf(w, result.Error())
	}
	//everything is ok
	fmt.Fprintf(w, "New user Inserted")
}

//to get all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//query the dao for all users
	users := services.GetAllUsers()
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
	//query the dao for update
	if update := dao.EditUser(currentUser.Username, edit); update {
		fmt.Fprintf(w, "User's informations successfully updated!")
	} else {
		fmt.Fprintf(w, "Failed to update username")
	}
}

//allow a user to change his password
func EditPasswordHandler(w http.ResponseWriter, r *http.Request) {

	//get old password and new password submitted by the user
	edit := struct {
		ActualPassword string `json:"actualpassword"`
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
	//query the dao to change the password
	if update := dao.EditPassword(hash, currentUser.Username); update {
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
	//call the dao to retrieve informations about the targeted user
	targetedUser, err := dao.GetUser(editor.Username)
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
	//everything is correct. Contact the dao to make the change
	if ok := dao.AdminEditUser(editor.Username, editor.User); !ok {
		fmt.Fprintf(w, "Failed to update user's infos!")
	} else {
		fmt.Fprintf(w, "User's infos updated!")
	}

}
