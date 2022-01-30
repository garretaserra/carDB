package CarService

import (
	"carDB/db"
	CarModel "carDB/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func ReadCar(rw http.ResponseWriter, r *http.Request) {
	readQuery := "SELECT id, brand, model, horse_power FROM car WHERE id = ?;"
	pathParams := mux.Vars(r)
	id := pathParams["id"]

	var resultCar CarModel.Car
	err := db.DBCon.QueryRow(readQuery, id).Scan(&resultCar.Id, &resultCar.Brand, &resultCar.Model, &resultCar.Horse_power)
	if err == sql.ErrNoRows {
		http.Error(rw, "404 car not found", http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("DB Error: ", err)
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(resultCar)
	if err != nil {
		fmt.Println("Marshal Error: ", err)
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}

	rw.Write(result)
}

// Initialize random source
var randomSource rand.Source = rand.NewSource(time.Now().UnixNano())
var rndMutex sync.Mutex

func CreateCar(rw http.ResponseWriter, r *http.Request) {
	var c CarModel.Car
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		fmt.Println("Parsing Error: ", err)
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}

	// Check all values exist
	if c.Brand == "" || c.Model == "" || c.Horse_power == 0 {
		http.Error(rw, "Format Error", http.StatusBadRequest)
		return
	}

	//Generate random id. Size of id should be much larger than stored entities to avoid id collisions
	rndMutex.Lock()
	c.Id = rand.New(randomSource).Intn(999999)
	rndMutex.Unlock()

	createQuery := "INSERT INTO car (id, brand, model, horse_power) VALUES (?, ?, ?, ?);"
	_, dbErr := db.DBCon.Exec(createQuery, c.Id, c.Brand, c.Model, c.Horse_power)
	if dbErr != nil {
		fmt.Println("DB Error: ", dbErr)
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Marshal Error: ", err)
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}

	rw.Write(result)
}

func DeleteCar(rw http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	id := pathParams["id"]

	deleteQuery := "DELETE FROM car WHERE id = ? RETURNING id, brand, model, horse_power;"
	var c CarModel.Car
	dbErr := db.DBCon.QueryRow(deleteQuery, id).Scan(&c.Id, &c.Brand, &c.Model, &c.Horse_power)
	if dbErr == sql.ErrNoRows {
		http.Error(rw, "404 car not found", http.StatusNotFound)
		return
	} else if dbErr != nil {
		fmt.Println("DB Error: ", dbErr)
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Marshal Error: ", err)
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(result)
}
