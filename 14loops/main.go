package main

import "fmt"

func main() {
	fmt.Println("Loops in Go Lang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for idx, i := range days {
		fmt.Printf("index is %v and value is %v\n", idx, i)
	}

	roughValue := 1
	for roughValue < 10 {
		if roughValue == 5 {
			break
		}
		fmt.Println("value is: ", roughValue)
		roughValue++
	}

	roughValue_1 := 20

	for roughValue_1 < 30 {
		if roughValue_1 == 29 {
			goto lco
		}
		fmt.Println("value is: ", roughValue_1)
		roughValue_1++
	}
lco:
	fmt.Println("Hello")
}
