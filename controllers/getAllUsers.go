package controllers

import (
	"fmt"
	"usermanager/dao"
	"usermanager/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	rows, err := dao.GetAllUsers()
	fmt.Println(rows)
	//close the rows once we read them all
	if err != nil {
		return users, err
	}
	defer rows.Close()
	//iterate through the rows
	for rows.Next() {
		var user models.User
		//store every column into the correct variable
		rows.Scan(&user.Firstname, &user.Lastname, &user.Username, &user.Role, &user.CreatedOn, &user.IsActive)
		users = append(users, user)
	}

	return users, nil

}
