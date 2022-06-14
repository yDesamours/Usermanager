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
	IsActive   bool          `json:"isactive"`
	CreatedOn  string        `json:"isActive"`
	ModifiedBy sql.NullInt64 `json:"modifiedBy"`
}

type ConnectString struct {
	Dbname   string
	Host     string
	Port     int
	Username string
	Password string
}
