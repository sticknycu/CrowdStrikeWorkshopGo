package main

import "fmt"
import "sync"

/**
Given an array of INTs, calculate the sum of even numbers in the array using 4 go routines in two different ways(using channels, using WaitGroup).
*/

var wg sync.WaitGroup
var mu sync.Mutex

func numberIsOdd(n int) bool {
	return n%2 == 0
}

func sum(v []int, c chan int) {
	s := 0
	for _, v := range v {
		if numberIsOdd(v) {
			s += v
		}
	}
	c <- s
}

var static_sum int

func sum2(v []int) int {
	total := 0
	for _, v := range v {
		if numberIsOdd(v) {
			total += v
		}
	}

	return total
}

func add(number int) {
	mu.Lock()
	static_sum += number
	mu.Unlock()

	wg.Done()
}

func main() {
	v := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println("Varianta cu Channels")

	c := make(chan int)

	go sum(v[:1/4*len(v)], c)
	go sum(v[1/4*len(v):2/4*len(v)], c)
	go sum(v[2/4*len(v):3/4*len(v)], c)
	go sum(v[3/4*len(v):], c)

	x1, x2, x3, x4 := <-c, <-c, <-c, <-c

	total := x1 + x2 + x3 + x4

	fmt.Println(total)

	fmt.Println("Varianta cu WaitGroup")

	total1 := sum2(v[:1/4*len(v)])
	total2 := sum2(v[1/4*len(v) : 2/4*len(v)])
	total3 := sum2(v[2/4*len(v) : 3/4*len(v)])
	total4 := sum2(v[3/4*len(v):])

	wg.Add(4)
	go add(total1)
	go add(total2)
	go add(total3)
	go add(total4)

	wg.Wait()

	fmt.Printf("Suma finala este: %d\n", static_sum)
}
