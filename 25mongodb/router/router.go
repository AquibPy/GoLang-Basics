package router

import (
	"github.com/AquibPy/mongoapi/controller"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler) // Serve Swagger UI

	return router
}
