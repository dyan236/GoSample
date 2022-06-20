package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	fmt.Println("Hello World")
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go incrementCounter(&waitGroup)
	}
	waitGroup.Wait()
	channel := make(chan int, 2)
	waitGroup.Add(2)
	go produce(channel, &waitGroup)
	go consume(channel, &waitGroup)
	waitGroup.Wait()
	list := LinkedList{
		head: nil,
	}
	for i := 0; i < 5; i++ {
		list.append(i)
	}
	list.printList()
	fmt.Println("Done")
}
