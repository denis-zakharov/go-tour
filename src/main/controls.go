package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// Accum demo
func Accum(n int) int {
	sum := 0
	// init and post statements are optional
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// ForWhile demo
func ForWhile(n int) int {
	sum := 1
	for sum < n {
		sum += sum
	}
	return sum
}

// ForeverLoop demo
func ForeverLoop(n int) {
	i := 0
	for {
		fmt.Println("Forever", i)
		i++
		if i > n {
			break
		}
	}
}

func pow(x, n, lim float64) float64 {
	result := lim
	if v := math.Pow(x, n); v < lim {
		result = v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return result
}

// Sqrt demo
func Sqrt(x float64, iters int) float64 {
	z := 1.0
	for i := 0; i < iters; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

// CheckOS switch demo
func CheckOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func incWeekday(d time.Weekday, i time.Weekday) time.Weekday {
	fmt.Println("Inc on", i)
	return d + i
}

// SatCheck switch demo laze eval
func SatCheck() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case incWeekday(today, 0):
		fmt.Println("Today.")
	case incWeekday(today, 1):
		fmt.Println("Tomorrow.")
	case incWeekday(today, 2):
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// SwitchTrue as an if-elif-else alternative
func SwitchTrue() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// DeferDemo demo
func DeferDemo() {
	fmt.Println("First stmt in the DeferDemo. Arg of the deferred call to Forever is immidiately evaluated.")
	defer ForeverLoop(int(Sqrt(2, 2)))
	fmt.Println("Last stmt in the DeferDemo")
}

// Countdown Multiple defers executed on a stack (LIFO queue)
func Countdown() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
