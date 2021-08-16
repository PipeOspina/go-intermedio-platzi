package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// go doSomething(i)
		go doSomethingSync(i, &wg)
	}

	wg.Wait()
}

func doSomething(i int) {
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Ended")
}

func doSomethingSync(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started Sync %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Ended Sync")
}
