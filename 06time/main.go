package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006")) // This date is standard

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020,time.August,11,23,23,0,0,time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))

}
