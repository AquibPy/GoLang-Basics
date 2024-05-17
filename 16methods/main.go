package main

import "fmt"

func main() {

	fmt.Println("Struct in golang")
	// no inheritance  and no super or parent in go lang
	aquib := User{"Aquib", "aquib@go.dev", true, 27}
	fmt.Println(aquib)

	fmt.Printf("Aquib Details are: %+v\n", aquib)

	fmt.Printf("Name is : %v and email is  %v\n", aquib.Name, aquib.Email)

	aquib.GetStatus()
	aquib.NewMail()

	fmt.Printf("Name is : %v and email is  %v\n", aquib.Name, aquib.Email)

	rect := Rectangle{25, 45}
	fmt.Println(rect)
	fmt.Println("Area is ", rect.Area())

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
