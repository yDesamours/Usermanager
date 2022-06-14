package dao

import (
	"usermanager/models"
)

//allows to change firstname, lastname, username
//accept a user structure as argument and the targeted username
func EditUser(username string, user models.User) error {
	db := GetDB()
	//set username to lowercase
	//execute the query. the exec function return the result of the query and an error
	result, err := db.Exec(updateUser, user.Firstname, user.Lastname, username)
	if err != nil {
		return err
	}
	//the affected method of the result return the number of rows affected by the query
	//if the number of rows is 0, the request failed
	if affected, err := result.RowsAffected(); affected == 0 {
		return err
	}

	return nil
}
