package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channel in golang")

	mych := make(chan int, 2)

	// mych <- 5
	// fmt.Println(<-mych)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	// recieve ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-mych
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-ch)
		// fmt.Println(<-ch)
		wg.Done()
	}(mych, wg)
	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		close(mych)
		wg.Done()
	}(mych, wg)

	wg.Wait()

}
