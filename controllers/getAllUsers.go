package controllers

import (
	"usermanager/dao"
	"usermanager/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	rows, err := dao.GetAllUsers()
	//close the rows once we read them all
	defer rows.Close()
	if err != nil {
		return users, err
	}
	//iterate through the rows
	for rows.Next() {
		var user models.User
		//store every column into the correct variable
		rows.Scan(&user.Firstname, &user.Lastname, &user.Username, &user.Role, &user.CreatedOn, &user.IsActive)
		users = append(users, user)
	}

	return users, nil

}
