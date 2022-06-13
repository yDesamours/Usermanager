package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"usermanager/models"

	_ "github.com/lib/pq"
)

//key for database connection
var credentials string

func createConnectionString() {
	var keys models.ConnectString
	//open json file
	keyFile, err := os.Open("database/keys.json")
	if err != nil {
		println(err.Error())
		panic("Can't open keys.json file!")
	}
	defer keyFile.Close()
	//read the file
	//io.ioutil return a slice of bytes that will be parse with the Unmarshal function of
	//encoding/json package
	content, _ := ioutil.ReadAll(keyFile)
	json.Unmarshal(content, &keys)

	//compose json string
	credentials = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", keys.Host, keys.Username, keys.Password, keys.Dbname)

}

//connection ctring required by sql.Open function
//var

// a pointer to store the sql.DB pointer return by sql.Open function
var DB *sql.DB

//function to initiate the connection
func Connect() {
	createConnectionString()
	db, err := sql.Open("postgres", credentials)
	if err != nil {
		panic("Cant connect to database!")
	}
	///test the connection
	err = db.Ping()
	if err != nil {
		panic("Can't reach database!")
	}
	//store the database object
	DB = db
}

//this function return the database object
func GetDB() *sql.DB {
	return DB
}
