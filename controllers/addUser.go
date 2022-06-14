package controllers

import (
	"usermanager/dao"
	"usermanager/models"
	"usermanager/utils"
)

func AddUser(newUser models.User) error {

	//test user credentials
	if err := utils.TestCredentials(newUser, true); err != nil {
		return err
	}

	//hash the submitted password
	newUser.Password = utils.HashPassword(newUser.Password)
	//lowercase all character, expect the password
	utils.Sanitize(&newUser)
	//insert thw user
	result := dao.InsertUser(newUser)

	return result
}
