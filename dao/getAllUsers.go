package dao

import (
	"database/sql"
	"fmt"
)

//function to get all users
func GetAllUsers() (*sql.Rows, error) {
	//get the database instance
	db := GetDB()
	//query the database. This statement return an error and all the rows selected by the query
	rows, err := db.Query(selectAllusers)
	//handle error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rows, nil
}
