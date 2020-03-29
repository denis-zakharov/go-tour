package util

import (
	"fmt"
	"math"
)

type abser interface {
	abs() float64
}

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/* Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type. */
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/* InterfaceDemo demo
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.
*/

//InterfaceDemo demo
func InterfaceDemo() {
	var a abser
	f := myFloat(-math.Sqrt2)
	describe(f)
	v := vector{3, 4}

	a = f // a MyFloat implements abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement abser.
	// Therefore should be a pointer for the assignment.
	a = &v
	describe(a)

	fmt.Println(a.abs())

	var nilVec *vector

	a = nilVec
	fmt.Println("Calling abs on a null pointer vector", a.abs())

	var any interface{}
	describe(any)

	typeAssertion()

	tryTypes(true)
	tryTypes("string")
	tryTypes(10)

	p1 := person{"Arthur Dent", 42}
	p2 := person{"Zaphod Beeblebrox", 9001}
	fmt.Println(p1, p2)
}

// cast
func typeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic without checking
	// f = i.(float64)
}

func tryTypes(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// Stringer interface (as toString)
type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
