package services

import (
	"usermanager/controllers"
	"usermanager/models"
)

func AddUser(newUser models.User) error {
	return controllers.AddUser(newUser)
}
