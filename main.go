package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// someRoutines()

	// someChannel()

	runSort()
}

func someRoutines() {
	f("direct")

	// A goroutine is a lightweight thread of execution
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func someChannel() {

	// Channels are the pipes that connect concurrent goroutines.
	// Used to send and recieve data from one go routine to another
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	// msg := <-messages
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
