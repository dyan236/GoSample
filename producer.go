package main

import "sync"

func produce(channel chan<- int, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
	waitGroup.Done()
}
