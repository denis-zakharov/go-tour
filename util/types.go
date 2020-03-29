package util

import (
	"fmt"
	"math"
	"strings"
)

type vertex struct {
	X int
	Y int
}

func pointersDemo() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(p) // see the address of a pointer
}

// StructDemo demo
func StructDemo() {
	v := vertex{Y: 2} // X defaults to 0
	fmt.Println(v)
	p := &v
	p.X = 1e9 // implicit dereference (*p).X
	fmt.Println(v, p)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// ArraysDemo demo
func ArraysDemo() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	sliceOfPrice1 := primes[0:4] // a dynamic size array
	sliceOfPrice2 := primes[3:]  // is like a pointer to a fixed array
	sliceOfPrice1[3] = 999
	fmt.Println(sliceOfPrice1)
	fmt.Println(sliceOfPrice2)
	fmt.Println(primes)

	sliceLiteral := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sliceLiteral)

	// capacity of an underlying array from the first element of a slice
	fmt.Printf("len=%d cap=%d %v\n", len(sliceLiteral[2:4]), cap(sliceLiteral[2:4]), sliceLiteral)

	// nil slice
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	// make slice of type with length (of an undelying zeroed array) and capacity
	aa := make([]int, 5)
	printSlice("aa", aa)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// append to d-slice beyond its capacity
	d = append(d, 1, 1, 1, 2, 2, 2)
	printSlice("d after append", d)

	// indexed for-each using range
	// other forms: for _, v; for i, _; for idx := range
	for i, v := range d {
		fmt.Printf("idx %d = %d\n", i, v)
	}
}

type coords struct {
	Lat, Long float64
}

var m map[string]coords // nil-map of str to coords struct

// MapsDemo demo
func MapsDemo() {
	fmt.Println(m)
	mp := make(map[string]coords)
	mp["key"] = coords{0, 42}
	fmt.Println(mp)

	var mapLiteral = map[string]coords{
		"Bell Labs": coords{
			40.68433, -74.39967,
		},
		"Google": coords{
			37.42202, -122.08408,
		},
	}
	fmt.Println(mapLiteral)

	mapLiteralShortcut := map[string]coords{
		"MyCompany": {123, 456},
	}
	fmt.Println(mapLiteralShortcut)
	myCompany, isPresent := mapLiteral["MyCompany"]
	if !isPresent {
		fmt.Println("There is no MyCompany in a mapLiteral", myCompany)
	}
	mapLiteral["MyCompany"] = mapLiteralShortcut["MyCompany"]
	myCompany, isPresent = mapLiteral["MyCompany"]
	if isPresent {
		fmt.Println("MyCompany is in the mapLiteral", myCompany)
	}
}

// WordCount demo
func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, v := range strings.Fields(s) {
		res[v]++
	}

	return res
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

// FuncValues demo
func FuncValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fibFunc := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibFunc())
	}

}
