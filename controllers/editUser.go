package controllers

import (
	"usermanager/dao"
	"usermanager/models"
	"usermanager/utils"
)

func EditUserControllers(edit, currentUser models.User) error {
	//thes for password matching. On failure, end the process
	if err := utils.ComparePassword(edit.Password, currentUser.Password); err != nil {
		return err
	}

	//test for credentials. On failure, end the proces
	if err := utils.TestCredentials(edit, false); err != nil {
		return err
	}
	//lowercase everything
	utils.Sanitize(&edit)
	//query the dao for update
	update := dao.EditUser(currentUser.Username, edit)

	return update
}
