package utils

import (
	"errors"
	"regexp"
	"strings"
	"usermanager/models"

	"golang.org/x/crypto/bcrypt"
)

func TestPassword(pass string) error {

	haveCapitalLetter := regexp.MustCompile("[A-Z]")
	haveLowerCaseLetter := regexp.MustCompile("[a-z]")
	haveNumber := regexp.MustCompile("[0-9]")
	if len(pass) < 8 {
		return errors.New("Password must have at least 8 characters")
	} else if !haveCapitalLetter.Match([]byte(pass)) {
		errors.New("Password must have at least 1 capital letter")
	} else if !haveLowerCaseLetter.Match([]byte(pass)) {
		errors.New("Password must have at least 1 lowercase letter")
	} else if !haveNumber.Match([]byte(pass)) {
		errors.New("Password must have at least 1 numeric character")
	}
	return nil
}

func HashPassword(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)

	return string(hash)
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func TestCredentials(userData models.User, pass bool) error {
	//test for firstname validity
	if len(userData.Firstname) == 0 {
		return errors.New("Must provide firstname")
	} else if len(userData.Firstname) > 50 {
		return errors.New("Too long firstname")
	}
	//test for lastname validity
	if len(userData.Lastname) == 0 {
		return errors.New("Must provide lastname")
	} else if len(userData.Lastname) > 50 {
		return errors.New("Too long lastname")
	}
	//tets for username validity
	if len(userData.Username) < 4 {
		return errors.New("Too short username")
	} else if len(userData.Username) > 50 {
		return errors.New("Too long username")
	}
	if pass {
		//test for password
		if err := TestPassword(userData.Password); err != nil {
			return errors.New("Mismatch password")
		}
	}
	return nil
}

func Sanitize(userData *models.User) {
	userData.Username = strings.ToLower(userData.Username)
}
