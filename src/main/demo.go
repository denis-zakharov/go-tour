package main

/* Basic Types~
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
import (
	"fmt"
	"math"
	"math/cmplx"
)

// bulk assignment demo
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// defaults
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	TypeConvDemo("Type conversions for constants")

	fmt.Printf("Accum 4: %d\n", Accum(4))
	fmt.Printf("ForWhile 1000: %d\n", ForWhile(1000))
	fmt.Println("Forever 2:")
	ForeverLoop(2)
	fmt.Println()
	fmt.Println("Return Sqrt after 5 iterations:", Sqrt(2, 5))
	fmt.Println("math.Sqrt", math.Sqrt(2))
	fmt.Println()
	CheckOS()
	fmt.Println()
	SatCheck()
	fmt.Println()
	SwitchTrue()
	fmt.Println()
	DeferDemo()
	fmt.Println()
	Countdown()
	fmt.Println()

	StructDemo()
	ArraysDemo()
	fmt.Println()

	MapsDemo()
	fmt.Println(WordCount("I ate a donut. When I ate another donut!"))

	fmt.Println()
	FuncValues()
	fmt.Println()

	MathDemo()
	fmt.Println()

	InterfaceDemo()

	ErrorDemo(0)
	ErrorDemo(1)

	Rot13Demo()
	fmt.Println()

	ParSumDemo()
	fmt.Println()

	FibStreamDemo()
	fmt.Println()

	FibSelectDemo()
	fmt.Println()

	DefaultSelectDemo()
	fmt.Println()

	WalkDemo()
	fmt.Println()

	MutexDemo()
	fmt.Println()

	CrawlDemo()
	fmt.Println()
}
