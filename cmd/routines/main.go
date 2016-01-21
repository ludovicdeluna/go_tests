package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	messageGoroutineA := make(chan bool, 1)
	messageGoroutineB := make(chan bool, 1)

	// Anonymous function into counter variable.
	counter := func(from string) {
		sleepTime := rand.Intn(10) * 100
		for i := 0; i < 3; i++ {
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			fmt.Println(from, ":", i)
		}
	}

	// Launch simple counter in my main "go routine"
	counter("direct")

	// Launch go routines in logical processor, pass counter as pointer
	go f("goroutine A", &counter, messageGoroutineA)
	go f("goroutine B", &counter, messageGoroutineB)

	// Show message to say : They are launched
	fmt.Println("All goroutines launched")

	// Wait for return of channel
	_, _ = <-messageGoroutineA, <-messageGoroutineB

	// Press a Key to end
	fmt.Println("Press a key to end program.")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func f(from string, counter *func(from string), done chan<- bool) {
	// Bad, isn't it ? But its because I ask myself if its possible to
	// send a pointer's function as callback. Answer is Yes :)
	// Luckily, I don't loop on a dereferencing pointer !
	(*counter)(from)
	done <- true
	close(done) // Optional. Need only if closed chan is checked (not the case here)
}
