package main

import "fmt"

func main() {

	fmt.Println("Struct in golang")
	// no inheritance  and no super or parent in go lang
	aquib := User{"Aquib", "aquib@go.dev", true, 27}
	fmt.Println(aquib)

	fmt.Printf("Aquib Details are: %+v\n", aquib)

	fmt.Printf("Name is : %v and email is  %v\n", aquib.Name, aquib.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
