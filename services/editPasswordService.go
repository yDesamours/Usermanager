package services

import (
	"io"
	"usermanager/controllers"
	"usermanager/models"
)

func EditpasswordService(r io.ReadCloser, user models.User) error {
	return controllers.EditPasswordController(r, user)
}
