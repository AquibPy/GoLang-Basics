package main

import (
    "fmt"
)

func sumOfDigits(n int) int {
    sum := 0
    for n != 0 {
        sum += n % 10 // Add the last digit to sum
        n /= 10       // Remove the last digit from n
    }
    return sum
}

func main() {
    num := 12345
    fmt.Println("Sum of digits:", sumOfDigits(num))
}
