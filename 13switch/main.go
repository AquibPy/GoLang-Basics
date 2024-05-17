package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and Case in Go Lang")

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	diceNumber := r.Intn(6) + 1
	fmt.Println("Value of Dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Printf("Dice value is %v and you can open\n", diceNumber)
	case 2:
		fmt.Printf("You can move %v spot\n", diceNumber)
	case 3:
		fmt.Printf("You can move %v spot\n", diceNumber)
	case 4:
		fmt.Printf("You can move %v spot\n", diceNumber)
	case 5:
		fmt.Printf("You can move %v spot\n", diceNumber)
	case 6:
		fmt.Printf("You can move %v spot and roll dice agian\n", diceNumber)
	default:
		fmt.Println("what was that ?")
	}
}

/**In Go, the fallthrough statement is used within a switch statement to transfer control
to the next case clause, even if the case expression does not match.**/
