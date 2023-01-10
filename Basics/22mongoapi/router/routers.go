package router

import (
	"github.com/gautamgc75/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// executes given view for specified url path
	router.HandleFunc("/api/movies" , controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie" , controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}" , controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}" , controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/delete-all-movie" , controller.DeleteAllMovies).Methods("DELETE")
	
	return router
}