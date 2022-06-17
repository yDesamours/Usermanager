package services

import (
	"usermanager/controllers"
	"usermanager/models"
)

func GetUser(username string) (models.UserResponse, error) {
	return controllers.GetUser(username)
}
