package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usermanager/dao"
	"usermanager/models"
	"usermanager/sessionHandlers"
	"usermanager/utils"
)

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
	if err := utils.ComparePassword(edit.ActualPassword, currentUser.Password); err != nil {
		fmt.Fprintf(w, "Incorect password")
		return
	}
	//test the new password
	if err := utils.TestPassword(edit.NewPassword); err != nil {
		fmt.Fprintf(w, err.Error())
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
	if err := utils.ComparePassword(editor.User.Password, currentUser.Password); err != nil {
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
