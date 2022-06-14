package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"usermanager/models"
)

//queries string
const (
	insertUser        = `INSERT INTO users (firstname, lastname, username, password, role_id) VALUES ($1, $2, $3, $4, $5)`
	selectAllusers    = `SELECT firstname, lastname, username, role_name, created_on, is_active FROM users`
	updateUser        = `UPDATE users SET firstname=$1, lastname=$2 WHERE username=$3`
	adminUpdateUser   = `UPDATE users SET firstname=$1, lastname=$2, username=$3, role=$4, is_active=$5 WHERE username=$6`
	updatePassword    = `UPDATE users SET password=$1 WHERE username=$2`
	desactivateUser   = `UPDATE users set is_active=false WHERE username=$1`
	getUser           = `SELECT firstname, lastname, username, role_name, is_active, created_on FROM users INNER JOIN role USING(role_id) WHERE username=$1`
	selectRoleID      = `SELECT role_id FROM roles where role_name=$1`
	getALLPermissions = `SELECT permission_name FROM roles INNER JOIN role_permission USING(permission_id) INNER JOIN roles USING(role_id) where role_name = $1`
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

//function to get all users
func GetAllUsers() (*sql.Rows, error) {
	//get the database instance
	db := GetDB()
	//query the database. This statement return an error and all the rows selected by the query
	rows, err := db.Query(selectAllusers)
	//handle error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//allows to change firstname, lastname, username
//accept a user structure as argument and the targeted username
func EditUser(username string, user models.User) error {
	db := GetDB()
	//set username to lowercase
	fmt.Println(user)
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

//allows to change a password
//accept the new password and the username as argument
func EditPassword(newPassword, username string) error {
	db := GetDB()

	//set username to lowewrcase
	result, err := db.Exec(updatePassword, newPassword, username)
	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); affected == 0 {
		return err
	}

	return nil
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
func GetUser(username string) (*models.User, error) {
	var savedUser models.User
	db := GetDB()
	var id int
	//set username to lowewrcase
	lowerName := strings.ToLower(username)
	//QueryRow return a single row
	err := db.QueryRow(getUser, lowerName).Scan(&id, &savedUser.Firstname, &savedUser.Lastname, &savedUser.Username, &savedUser.Role, &savedUser.IsActive, &savedUser.CreatedOn)

	if err != nil {
		return nil, err
	}

	return &savedUser, nil
}

//this function return the id of a role
func GetRoleID(role_name string) string {
	db := GetDB()
	var role_ID string
	err := db.QueryRow(selectRoleID, role_name).Scan(&role_ID)
	if err != nil {
		return "role doesn't exists"
	}

	return role_ID
}

//this function return all the permissions attached to a role
func GetRolePermissions(roleName string) ([]string, error) {
	db := GetDB()
	var permissions []string
	rows, err := db.Query(getALLPermissions, roleName)
	defer rows.Close()

	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		var permission string

		rows.Scan(&permission)
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
