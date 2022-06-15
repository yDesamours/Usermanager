package dao

import (
	"errors"
	"usermanager/models"
)

//allows and admin to change all the user info except the the password
//accept the actual username as argument and the new user infos
func AdminEditUser(username string, user models.User) error {
	db := GetDB()

	result, err := db.Exec(adminUpdateUser, user.Firstname, user.Lastname, user.Username, user.Role, user.IsActive, username)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New("No row affected!")
	}

	return nil
}
