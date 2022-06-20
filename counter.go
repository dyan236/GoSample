package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mutex = sync.RWMutex{}

func incrementCounter(waitGroup *sync.WaitGroup) {
	mutex.Lock()
	counter++
	fmt.Printf("Counter value: %v\n", counter)
	mutex.Unlock()
	waitGroup.Done()
}
