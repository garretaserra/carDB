package main

import (
	"carDB/db"
	CarRouter "carDB/routing"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	// Read CLI flags
	env := flag.String("env", "local", "Indicate which environment youa re working in local or dev")
	flag.Parse()

	// Load environement
	var err error
	switch *env {
	case "local":
		err = godotenv.Load("./envs/local.env")
	case "dev":
		err = godotenv.Load("./envs/dev.env")
	default:
		err = godotenv.Load("./envs/local.env")
	}

	if err != nil {
		log.Fatal("Error loading .env file", err, *env)
	}

	// Open connection to SQL
	db.OpenSQLConnection()
	defer db.DBCon.Close()

	srv := mux.NewRouter()

	// Add car router
	CarRouter.AddCarRoutes(srv.PathPrefix("/cars").Subrouter())

	address := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Println("Starting server at " + address)
	log.Fatal(http.ListenAndServe(address, srv))
}
