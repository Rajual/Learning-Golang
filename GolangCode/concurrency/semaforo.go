package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id %d started", i)
	fmt.Println("")
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished", i)
	fmt.Println("")
	<-c
}
