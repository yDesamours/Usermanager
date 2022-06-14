package services

import (
	"io"
	"usermanager/controllers"
	"usermanager/models"
)

func AdminUpdateUser(r io.ReadCloser, user models.User) error {
	return controllers.AdminEditUserController(r, user)
}
