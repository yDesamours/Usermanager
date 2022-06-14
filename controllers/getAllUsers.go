package controllers

import (
	"usermanager/dao"
	"usermanager/models"
)

func GetAllUsers() []models.User {
	return dao.GetAllUsers()
}
