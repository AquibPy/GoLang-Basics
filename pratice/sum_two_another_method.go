package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza Rating")
	fmt.Println("Please rate pizza between 1-5")

	// Reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Prompt and read first rating
	fmt.Print("Enter the first rating: ")
	input1, _ := reader.ReadString('\n')

	// Prompt and read second rating
	fmt.Print("Enter the second rating: ")
	input2, _ := reader.ReadString('\n')

	// Trim leading and trailing whitespaces, then convert to float64
	rating1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println("Error parsing first rating:", err)
		return
	}

	rating2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println("Error parsing second rating:", err)
		return
	}

	// Calculate the sum of ratings
	sum := rating1 + rating2

	// Print the sum
	fmt.Printf("The sum of %.2f and %.2f is %.2f\n", rating1, rating2, sum)
}
