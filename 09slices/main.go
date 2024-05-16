package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777

	fmt.Println(highScore)

	highScore = append(highScore, 555, 666, 321)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println("Sorted:", highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	fmt.Println(highScore)

	// remove value from slices based on index

	var courses = []string{"python", "Ruby", "JS", "C++", "C"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
