package functions

import (
	"fmt"
	"time"
)

// anon func's are not very recommended for DRY purposes. Only use it when it is necessary.
func anon() {
	x := 5

	// Anonymous func 1
	y := func(x int) int {
		return x * 2
	}
	fmt.Println("y anon func: ", y(x))

	// Anonymous func 2 (it calls directly the param and calls itself)
	z := func() int {
		return x * 2
	}()
	fmt.Println("y anon func: ", z)

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

///////////////////////

/* Variadic funcs */
// Variadic func able us to insert a dinamic amount of params
// The param is a slice of the defined type
func sum(numbers ...int) int {
	totalSum := 0
	for _, num := range numbers {
		totalSum += num
	}

	return totalSum
}

func getNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

///////////////////////
/* Named returns func's */
func multiples(a int8) (double, triple, quad int8) {
	double = a * 2
	triple = a * 3
	quad = a * 4
	return
}

func Functions() {
	anon()

	fmt.Println("----------------")
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	fmt.Println("----------------")
	getNames("Nico", "Pedro")
	getNames("Nico", "Pedro", "Facu")

	fmt.Println("----------------")
	double, triple, quad := multiples(2)
	// %o is for base 8 int
	fmt.Printf("Double: %o, triple: %o, quad: %o ", double, triple, quad)

}
