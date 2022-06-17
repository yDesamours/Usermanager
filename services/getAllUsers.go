package services

import (
	"usermanager/controllers"
	"usermanager/models"
)

func GetAllUsers() ([]models.UserResponse, error) {
	return controllers.GetAllUsers()
}
