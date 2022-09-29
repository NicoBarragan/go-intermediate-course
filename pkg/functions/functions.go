package functions

import (
	"fmt"
	"time"
)

func lambda() {
	x := 5

	// Anonymous func 1
	y := func(x int) int {
		return x * 2
	}
	fmt.Println("y lambda func: ", y(x))

	// Anonymous func 2 (it calls directly the param and calls itself)
	z := func() int {
		return x * 2
	}()
	fmt.Println("y lambda func: ", z)

	// 3rd Anonymous way: lambda Goroutine in a channel
	// Create channel for blocking Gorouting until msg arrives
	c := make(chan int)
	// anon Goroutine
	go func() {
		// Exec func logic in Goroutine
		fmt.Println("Starting Goroutine anon func...")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine ended")

		// Send int to c channel
		c <- 1
	}() // calls itself
	// Waits until receiving the msg in c channel
	<-c

}

func Functions() {
	lambda()
}
