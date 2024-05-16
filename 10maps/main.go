package main

import "fmt"

func main() {

	fmt.Println("Maps in Go Lang")
	languages := make(map[string]string)

	languages["Python"] = "Data Science"
	languages["JS"] = "Web Dev"
	languages["Java"] = "Backend"

	fmt.Println("List of all langauges:", languages)
	fmt.Println("Python", languages["Python"])

	delete(languages, "Java")

	fmt.Println("List of all languages:", languages)

	// loops

	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}

	
}
