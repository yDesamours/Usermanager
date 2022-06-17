package controllers

import (
	"bytes"
	"encoding/gob"
	"usermanager/dao"
	"usermanager/models"
)

func GetUser(username string) (models.UserResponse, error) {
	user, err := dao.GetUser(username)
	var response models.UserResponse
	if err != nil {
		return response, err
	}

	var buffer bytes.Buffer

	gob.NewEncoder(&buffer).Encode(user)
	gob.NewDecoder(&buffer).Decode(&response)

	return response, nil
}
