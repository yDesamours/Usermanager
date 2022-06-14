package dao

import "strings"

func ErrorHandler(err error) string {

	switch {
	case strings.Contains(err.Error(), "duplicate"):
		return "Username already exists! Try another."
	case strings.Contains(err.Error(), "no rows in result set"):
		return "User does'nt exist"
	}
	return ""
}
