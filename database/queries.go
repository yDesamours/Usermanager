package database

import (
	"fmt"
	"strings"
	"usermanager/models"
)

//queries string
const (
	insertUser      = `INSERT INTO users (firstname, lastname, username, password) VALUES ($1, $2, $3, $4)`
	selectAllusers  = `SELECT firstname, lastname, username, role From users`
	updateUser      = `UPDATE users SET firstname=$1, lastname=$2, username=$3 WHERE username=$4`
	adminUpdateUser = `UPDATE users SET firstname=$1, lastname=$2, username=$3, role=$4, is_active=$5 WHERE username=$6`
	updatePassword  = `UPDATE users SET password=$1 WHERE username=$2`
	desactivateUser = `UPDATE users set is_active=false WHERE username=$1`
	getUser         = `SELECT * FROM users WHERE username=$1`
)

//Function for inserting a new user
//accept a user structure as argument
func InsertUser(newUser models.User) string {
	//get the db instance
	db := GetDB()
	//query the database
	_, err := db.Exec(insertUser, newUser.Firstname, newUser.Lastname, newUser.Username, newUser.Password)
	//if there is an error, handle it
	if err != nil {
		response := ErrorHandler(err)
		return response
	}
	//insertion succed
	return "New user " + newUser.Username + " inserted!\nTry to connect."
}

//function to get all users
func GetAllUsers() []models.User {
	//get the database instance
	db := GetDB()
	//all the user will be store into a slice
	var users []models.User
	//query the database. This statement return an error and all the rows selected by the query
	rows, err := db.Query(selectAllusers)
	//handle error
	if err != nil {
		fmt.Println(err)
	}
	//close the rows once we read them all
	defer rows.Close()
	//iterate through the rows
	for rows.Next() {
		var user models.User
		//store every column into the correct variable
		rows.Scan(&user.Firstname, &user.Lastname, &user.Username, &user.Role)
		users = append(users, user)
	}

	return users
}

//allows to change firstname, lastname, username
//accept a user structure as argument and the targeted username
func EditUser(username string, user models.User) bool {
	db := GetDB()
	//set username to lowercase
	fmt.Println(user)
	//execute the query. the exec function return the result of the query and an error
	result, err := db.Exec(updateUser, user.Firstname, user.Lastname, user.Username, username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	//the affected method of the result return the number of rows affected by the query
	//if the number of rows is 0, the request failed
	if affected, _ := result.RowsAffected(); affected == 0 {
		return false
	}

	return true
}

//allows to change a password
//accept the new password and the username as argument
func EditPassword(newPassword, username string) bool {
	db := GetDB()

	//set username to lowewrcase
	result, err := db.Exec(updatePassword, newPassword, username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return false
	}

	return true
}

//allows and admin to change all the user info except the the password
//accept the actual username as argument and the new user infos
func AdminEditUser(username string, user models.User) bool {
	db := GetDB()

	result, err := db.Exec(adminUpdateUser, user.Firstname, user.Lastname, user.Username, user.Role, user.IsActive, username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return false
	}

	return true
}

//return all the user info correspoding to the username passed as parameter
func GetUser(username string) (models.User, string) {
	var savedUser models.User
	db := GetDB()
	var id int
	//set username to lowewrcase
	lowerName := strings.ToLower(username)
	//QueryRow return a single row
	err := db.QueryRow(getUser, lowerName).Scan(&id, &savedUser.Firstname, &savedUser.Lastname, &savedUser.Username, &savedUser.IsActive, &savedUser.Password, &savedUser.Role)

	if err != nil {
		errorMessage := ErrorHandler(err)
		return savedUser, errorMessage
	}

	return savedUser, ""
}
