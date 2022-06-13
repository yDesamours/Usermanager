package models

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	IsActive  bool   `json:"isactive"`
}

type ConnectString struct {
	Dbname   string
	Host     string
	Port     int
	Username string
	Password string
}
