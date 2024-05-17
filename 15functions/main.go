package main

import "fmt"

func main() {
	fmt.Println("Functions in go lang")
	greet()
	fmt.Println("Sum of two number : ", sum(2, 5))

	fmt.Println("Pro Added : ", proAdder(2, 5, 8, 7))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Pro Added : ", proAdder(nums...))

	age, name := greeting("Aquib")
	fmt.Println(age, name)

}

func greet() {
	fmt.Println("Hello from GoLang")
}

func sum(num1 int, num2 int) int {
	out := num1 + num2
	return out

}

func proAdder(values ...int) int {
	total := 0

	for _, i := range values {
		total += i
	}
	return total
}

func greeting(name string) (int, string) {

	nam := "Mohd " + name
	age := 27
	return age, nam
}
