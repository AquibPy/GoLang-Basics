package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	sum := num1 + num2

	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
}
