package main

import (
	"fmt"
	"sync"
)

var v map[string]int
var wg sync.WaitGroup
var mu sync.Mutex

func inc(key string) {
	mu.Lock()
	v[key]++
	mu.Unlock()

	wg.Done()
}

func getValue(key string) int {
	return v[key]
}

func main() {
	println("---- MAPS ----")

	v = map[string]int{}

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go inc("mykey")
	}

	wg.Wait()

	fmt.Println(getValue("mykey"))
}
