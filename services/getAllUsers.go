package services

import (
	"usermanager/controllers"
	"usermanager/models"
)

func GetAllUsers() []models.User {
	return controllers.GetAllUsers()
}
