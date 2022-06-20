package main

import (
	"fmt"
	"sync"
)

func consume(channel <-chan int, waitGroup *sync.WaitGroup) {
	for i := range channel {
		fmt.Printf("Consumer output: %v\n", i)
	}
	waitGroup.Done()
}
