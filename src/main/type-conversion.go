package main

import (
	"fmt"
	"math"
)

// constant may have context type conversion
const (
	Big   = 1 << 100
	Small = 1 << 1
)

func needInt(x int) int {
	return x * 2
}

func needFloat(x float64) float64 {
	return x * 0.001
}

// TypeConvDemo exported
func TypeConvDemo(msg string) {
	fmt.Printf("\n === %s ===\n", msg)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	var myf float64 = float64(x)
	fmt.Println(x, y, z, myf)

	// conversion by context
	fmt.Println("Small", needInt(Small))
	fmt.Println("Small float", needFloat(Small))
	fmt.Println("Big", needFloat(Big))

	fmt.Println()
}
