package main

import "fmt"

// A variable is said to be shadowing another variable if it “overrides” the variable in a more specific scope.
var var_shadow int = 32

// jwttoken := 300000  outside the method walrus operator is not allowed

const LoginToken string = "gssadcsfdd"

var (
	name       string = "Aquib"
	company    string = "Extreme networks"
	experience int    = 2
)

func main() {

	var user string = "aquib"
	fmt.Println((user))
	fmt.Printf("Variable %v is of type: %T \n", user, user)

	var isbool bool = true
	fmt.Println((isbool))
	fmt.Printf("Variable %v is of type: %T \n", isbool, isbool)

	var smallFloat float32 = 255.455495959
	fmt.Println((smallFloat)) //output : 255.45549
	fmt.Printf("Variable %v is of type: %T \n", smallFloat, smallFloat)

	var largeFloat float64 = 255.455495959
	fmt.Println((largeFloat)) // here we will get same value
	fmt.Printf("Variable %v is of type: %T \n", largeFloat, largeFloat)

	// default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar) // output - 0
	fmt.Printf("Variable %v is of type: %T \n", anotherVar, anotherVar)

	// implicit type
	var username = "AquibPy"
	fmt.Println(username)

	// no var style
	num_of_user := 30000
	fmt.Println(num_of_user)

	fmt.Println(LoginToken)

	fmt.Println(name, company, experience)

	//variable shadowing
	fmt.Println("Before", var_shadow)
	var_shadow := 32.658996
	fmt.Println("After", var_shadow)
}
