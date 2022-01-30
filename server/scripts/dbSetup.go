package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/car?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Read contents of startup.SQL
	dat, err := os.ReadFile("./scripts/startup.SQL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Executing: \n", string(dat))

	// Connect and execute SQL start up
	_, err = db.Exec(string(dat))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished executing script")
}
