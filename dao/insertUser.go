package dao

import (
	"usermanager/models"
)

//Function for inserting a new user
//accept a user structure as argument
func InsertUser(newUser models.User) error {
	//get the db instance
	db := GetDB()
	//get the role id
	role_ID := GetRoleID(newUser.Role)
	//query the database
	_, err := db.Exec(insertUser, newUser.Firstname, newUser.Lastname, newUser.Username, newUser.Password, role_ID)
	//if there is an error, handle it
	if err != nil {
		//response := ErrorHandler(err)
		return err
	}
	//insertion succed
	return nil
}
