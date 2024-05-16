package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)
	mynum := 23

	var ptr = &mynum
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value is", mynum)
}
