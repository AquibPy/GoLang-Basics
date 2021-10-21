//  A variable is said to be shadowing another variable if it “overrides” the variable in a more specific scope.

package main

import (
	"fmt"
)

var i int = 22

func main() {
	fmt.Println(i)
	var i float32 = 22.3
	fmt.Print(i)
}
