package main

import "fmt"

func main() {
	fmt.Println("Max of two number")

	var num1, num2 int

	fmt.Print("Enter First number:")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Print("Enter Second number:")
	_, err_1 := fmt.Scan(&num2)
	if err_1 != nil {
		fmt.Println("Error reading input:", err_1)
		return
	}

	if num1 > num2 {
		fmt.Printf("%v is greater than %v\n", num1, num2)
	} else if num1 == num2 {
		fmt.Printf("Numbers are same\n")
	} else {
		fmt.Printf("%v is greater than %v\n", num2, num1)
	}

	if num1%2 == 0 {
		fmt.Printf("%v is even number\n", num1)
	} else {
		fmt.Printf("%v is odd number\n", num1)
	}

}
