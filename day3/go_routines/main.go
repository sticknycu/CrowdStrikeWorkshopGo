package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	go say("hello")
	say("world")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	println("--- Fibonacci ---")

	go fibonacci(10, c)
	//d := <-c
	// print(d)

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("--- functii anonime ---")

	go func(s string, x int) {
		fmt.Printf("Am string-ul %s si int-ul %d\n", s, x)
	}("Ana", 15)

	cc := make(chan int)
	quit := make(chan error)

	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(<-cc)
		}
		quit <- fmt.Errorf("finished loop")
	}()

	r, rt := 0, 1
	for {
		select {
		case cc <- r:
			r, rt = rt, r+rt
		case err := <-quit:
			fmt.Println("quit", err)
			return
		}
	}
}
