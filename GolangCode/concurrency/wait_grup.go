package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	counter0 := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
	fmt.Println(time.Since(counter0))
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started %d", i)
	fmt.Println("")
	time.Sleep(4 * time.Second)
	fmt.Println("End")
}
