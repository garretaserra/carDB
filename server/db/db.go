package db

import (
	"database/sql"
	"fmt"
	"os"
)

// Used to hold the database connection handle
var DBCon *sql.DB

func OpenSQLConnection() {
	sqlConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))

	var err error
	DBCon, err = sql.Open("mysql", sqlConnectionString)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
