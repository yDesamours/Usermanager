package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"usermanager/dao"
	"usermanager/models"
	"usermanager/utils"
)

func EditPasswordController(r io.ReadCloser, user models.User) error {
	//get old password and new password submitted by the user
	edit := struct {
		ActualPassword string `json:"actualpassword"`
		NewPassword    string `json:"newpassword"`
	}{}
	json.NewDecoder(r).Decode(&edit)

	//test for password matching
	if ok := utils.ComparePassword(edit.ActualPassword, user.Password); !ok {
		return errors.New("Incorrect passowrd")
	}
	//test the new password
	if err := utils.TestPassword(edit.NewPassword); err != nil {
		return err
	}
	//hash the password
	hash := utils.HashPassword(edit.NewPassword)

	update := dao.EditPassword(hash, user.Username)

	return update
}
