package util

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

// ParSumDemo demo
func ParSumDemo() {
	s := []int{7, 2, 8, -9, 4, 0}

	optionalBufferSize := len(s)
	c := make(chan int, optionalBufferSize) // channel

	// goroutines
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func fibStream(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // close channel
}

// FibStreamDemo demo
func FibStreamDemo() {
	c := make(chan int, 100)
	go fibStream(cap(c), c)
	for i := range c { // v, ok := <-ch, ok is false for a closed channel
		fmt.Println(i)
	}
}

func fibSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // wait on each case until any can be run
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// FibSelectDemo demo
func FibSelectDemo() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibSelect(c, quit)
}

//DefaultSelectDemo demo
func DefaultSelectDemo() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
