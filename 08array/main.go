package main

import "fmt"

func main() {

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"
	// fruitList[0] = "Apple"
	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list length:", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Cabbage"}
	fmt.Println("Fruit list is:", vegList)
	fmt.Println("Fruit list length:", len(vegList))

}
