package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"usermanager/dao"
	"usermanager/models"
	"usermanager/utils"
)

func AdminEditUserController(r io.ReadCloser, currentUser models.User) error {

	//etract the json data from the request. Copy them into a structure
	//This structure has 2 fields. One for storing new data about the user, but the password sent his the admin's
	//password. The other fields hold the username of the user whose data must be updated
	editor := struct {
		User     models.User `json:"user"`
		Username string      `json:"targetusername"`
	}{}
	json.NewDecoder(r).Decode(&editor) //all data are strored.fmt.Println(editor)
	//call the dao to retrieve informations about the targeted user

	targetedUser, err := dao.GetUser(editor.Username)
	if err != nil {
		return err
	}
	//If the targeted user is an admin, stop the process
	if targetedUser.Role == "admin" {
		return errors.New("You can't edit another admin's informations")
	}

	//check for credentials
	//If the test failed, stop the process and return a message describing what's wrong
	if err := utils.TestCredentials(editor.User, false); err != nil {
		return err
	}
	//test the password of the user
	//if the test failed, what do we do? Stop the process. On connait la chanson
	if ok := utils.ComparePassword(editor.User.Password, currentUser.Password); !ok {
		return errors.New("Incorect password")

	}

	if updated := dao.AdminEditUser(editor.Username, editor.User, currentUser.Username); updated != nil {
		return updated
	}
	return nil
}
