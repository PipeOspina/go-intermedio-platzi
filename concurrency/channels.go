package main

import "fmt"

func main() {
	// Unbuferred Channel (no executable in main it needs concurrent functions to send value)
	/*
		c := make(chan int)

		c <- 1

		fmt.Println(<-c)
	*/

	// Buferred Channel

	c := make(chan int, 4)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
