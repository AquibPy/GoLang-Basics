package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AquibPy/mongoapi/router"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/AquibPy/mongoapi/docs" // Important: This imports the generated docs
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler) // Serve Swagger UI

	fmt.Println("Server is getting started....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}