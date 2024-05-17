package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int16 = 25
	fmt.Printf("%v, %T", i, i)

	fmt.Println()

	var j float32
	j = float32(i)

	fmt.Printf("%v, %T", j, j)

	fmt.Println()

	var k string
	k = strconv.Itoa(int(i))

	fmt.Printf("%v, %T", k, k)
	fmt.Println()

	var t string
	t = fmt.Sprintf("%f", j)

	fmt.Printf("%v, %T", t, t)

}
