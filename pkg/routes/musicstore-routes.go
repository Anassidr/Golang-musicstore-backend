package routes

import (
	"github.com/anassidr/go-musicstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMusicStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/instrument/", controllers.CreateInstrument).Methods("POST")
	router.HandleFunc("/instrument/", controllers.GetInstrument).Methods("GET")
	router.HandleFunc("/instrument/{instrumentId}", controllers.GetInstrumentById).Methods("GET")
	router.HandleFunc("/instrument/{instrumentId}", controllers.UpdateInstrument).Methods("PUT")
	router.HandleFunc("/instrument/{instrumentId}", controllers.DeleteInstrument).Methods("DELETE")
}
