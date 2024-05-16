package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("welcome to maths in Go Lang")

	// random number
	// fmt.Println(rand.Intn(5))

	// random from crypto
	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println((myRandNum))
}
