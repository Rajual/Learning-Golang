package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exists := m.cache[key]
	m.lock.Unlock()
	if !exists {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}

	return result.value, result.err
}

//interface{} is same a <T> on Dart
func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cacle := NewCache(GetFibonacci)
	fibo := []int{42, 41, 42, 38}
	var wg sync.WaitGroup
	for _, v := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			start := time.Now()

			value, err := cacle.Get(index)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("%d, %s, %d", index, time.Since(start), value)
			fmt.Println("")
		}(v)

	}
	wg.Wait()
}
