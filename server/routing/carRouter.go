package CarRouter

import (
	CarService "carDB/services"

	"github.com/gorilla/mux"
)

func AddCarRoutes(srv *mux.Router) {
	srv.Methods("GET").Path("/{id}").HandlerFunc(CarService.ReadCar)
	srv.Methods("POST").Path("").HandlerFunc(CarService.CreateCar)
	srv.Methods("DELETE").Path("/{id}").HandlerFunc(CarService.DeleteCar)
}
