package services

import (
	"usermanager/controllers"
	"usermanager/models"
)

func GetAllUsers() ([]models.User, error) {
	return controllers.GetAllUsers()
}
