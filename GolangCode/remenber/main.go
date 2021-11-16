package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x, y)
	myValue, err := strconv.ParseInt("p", 0, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	s := []int{1, 2, 3}

	for index, value := range s {
		fmt.Println(index, value)
	}

	s = append(s, 12)

	c := make(chan int)

	go doSomething(c)

	<-c

}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Print("Done")
	c <- 1
}
