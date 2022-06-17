package models

import (
	"database/sql"
)

type User struct {
	Firstname  string        `json:"firstname"`
	Lastname   string        `json:"lastname"`
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	Role       string        `json:"role"`
	IsActive   bool          `json:"isActive"`
	CreatedOn  string        `json:"createdOn"`
	ModifiedBy sql.NullInt64 `json:"modifiedBy"`
}

type UserResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	IsActive  bool   `json:"isActive"`
	CreatedOn string `json:"createdOn"`
}
