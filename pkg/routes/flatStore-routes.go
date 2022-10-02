package routes

import (
	"example/rental_app/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterFlatStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/flat/", controllers.CreateFlat).Methods("POST")
	router.HandleFunc("/flat/", controllers.GetFlat).Methods("GET")
	router.HandleFunc("/flat/{flatId}", controllers.GetFlatById).Methods("GET")
	router.HandleFunc("/flat/{flatId}", controllers.UpdateFlat).Methods("PUT")
	router.HandleFunc("/flat/{flatId}", controllers.DeleteFlat).Methods("DELETE")
}
