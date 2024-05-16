package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Rating:")

	//comma ok syntax or error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating is", input)
	fmt.Printf("Type of this rating is %T \n", input)
	/**
	Here we are giving rating as int but %T will tell us that input is string
	**/

	
}
