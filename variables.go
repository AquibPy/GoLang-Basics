package main

import (
	"fmt"
)

var m int16 = 12 // outside the main func we can't declare variables using :

var (
	name   string = "Aquib"
	course string = "MCA"
	year   int    = 3
)

func main() {
	// variable declaration
	var i int
	i = 42

	var j int = 53

	k := 55

	var l float32 = 99

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	fmt.Printf("%v, %T", l, l) // This will print 99, float32
	fmt.Println()
	fmt.Printf("%v, %T", m, m)

	fmt.Println()

	fmt.Println(name, course, year)
}
