package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("web request in go lang")

	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Reponse type %T\n", response) // output: Reponse type *http.Response

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println("Reponse is ", string(databytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
