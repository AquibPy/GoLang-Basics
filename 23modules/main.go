package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in Go lang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal((http.ListenAndServe(":4000", r)))

}

func greeter() {
	fmt.Println("Het there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome Go Lang</h1>"))
}
