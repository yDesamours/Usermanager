package services

import (
	"usermanager/controllers"
	"usermanager/models"
)

func EditUserService(edit, currentUser models.User) error {
	return controllers.EditUserControllers(edit, currentUser)
}
