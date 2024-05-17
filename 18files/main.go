package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in GO")
	content := "This need to go in file"
	files, err := os.Create("./files.txt")
	checkNilErr(err)
	length, err := io.WriteString(files, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer files.Close()

	readFile("./files.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
