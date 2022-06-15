package dao

import (
	"strings"
	"usermanager/models"
)

//return all the user info correspoding to the username passed as parameter
func GetUser(username string) (*models.User, error) {
	var savedUser models.User
	db := GetDB()
	//set username to lowewrcase
	lowerName := strings.ToLower(username)
	//QueryRow return a single row
	err := db.QueryRow(getUser, lowerName).Scan(&savedUser.Firstname, &savedUser.Lastname, &savedUser.Username, &savedUser.Role, &savedUser.IsActive, &savedUser.CreatedOn, &savedUser.Password)

	if err != nil {
		return nil, err
	}

	return &savedUser, nil
}
