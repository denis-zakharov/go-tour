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

	"github.com/denis-zakharov/go-tour/util"
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

	util.TypeConvDemo("Type conversions for constants")

	fmt.Printf("Accum 4: %d\n", util.Accum(4))
	fmt.Printf("ForWhile 1000: %d\n", util.ForWhile(1000))
	fmt.Println("Forever 2:")
	util.ForeverLoop(2)
	fmt.Println()
	fmt.Println("Return Sqrt after 5 iterations:", util.Sqrt(2, 5))
	fmt.Println("math.Sqrt", math.Sqrt(2))
	fmt.Println()
	util.CheckOS()
	fmt.Println()
	util.SatCheck()
	fmt.Println()
	util.SwitchTrue()
	fmt.Println()
	util.DeferDemo()
	fmt.Println()
	util.Countdown()
	fmt.Println()

	util.StructDemo()
	util.ArraysDemo()
	fmt.Println()

	util.MapsDemo()
	fmt.Println(util.WordCount("I ate a donut. When I ate another donut!"))

	fmt.Println()
	util.FuncValues()
	fmt.Println()

	util.MethodsDemo()
	fmt.Println()

	util.InterfaceDemo()

	util.ErrorDemo(0)
	util.ErrorDemo(1)

	util.Rot13Demo()
	fmt.Println()

	util.ParSumDemo()
	fmt.Println()

	util.FibStreamDemo()
	fmt.Println()

	util.FibSelectDemo()
	fmt.Println()

	util.DefaultSelectDemo()
	fmt.Println()

	util.WalkDemo()
	fmt.Println()

	util.MutexDemo()
	fmt.Println()

	util.CrawlDemo()
	fmt.Println()
}
