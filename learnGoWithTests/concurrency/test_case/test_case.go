package main

import (
	"fmt"
	"sync"
)

type result struct {
	string
	bool
}

func main() {
	var wg sync.WaitGroup
	resultChannel := make(chan result)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: trying to send\n", i)
			resultChannel <- result{fmt.Sprintf("url-%d", i), true}
			fmt.Printf("Goroutine %d: send complete\n", i)
		}(i)
	}

	go func() {
		for i := 0; i < 3; i++ {
			r := <-resultChannel
			fmt.Println("Main received:", r)
		}
	}()

	wg.Wait()
}
