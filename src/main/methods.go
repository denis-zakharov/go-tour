package main

import (
	"fmt"
	"math"
)

type vector struct {
	X, Y float64
}

func (v *vector) abs() float64 {
	if v == nil {
		return -1
	}
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// receiver should be a pinter to modify itself
// otherwise scale will get a copy
func (v *vector) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// MathDemo demo: auto-(de)referencing works only for methods.
// Arguments of regular functions should be of explicit types.
func MathDemo() {
	vec := vector{3, 4}
	fmt.Println(vec.abs())
	vec.scale(10) // auto-ref: (&vec)
	fmt.Println(vec.abs())
	fmt.Println((&vec).abs()) // auto-deref: *(&vec) - a value of a vector pointer
}
