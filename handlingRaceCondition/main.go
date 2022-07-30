package main

import (
	"fmt"
	"sync"
)

// Hit object
type Hit struct {
	count int
	sync.Mutex
}

// for goroutine operations
var wg sync.WaitGroup

func main() {
	// create a hit object
	hit := &Hit{}
	// result needed for validation
	count := 10000

	for i := 0; i < count; i++ {
		wg.Add(1)

		// create an anonymous function, and add go keyword for goroutines
		go func(hit *Hit) {

			hit.Lock() // access block

			hit.count++ // increments the number

			hit.Unlock() // unblock access

			wg.Done()

		}(hit)
	}

	wg.Wait()
	fmt.Println("Must be", count, "\nWhat we have ", hit.count, "\nLoss", count-hit.count)
}
