package util

import (
	"fmt"
	"time"
)

type myError struct {
	When time.Time
	What string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run(i int) (string, error) {
	if i == 0 {
		return "FAIL", &myError{
			time.Now(),
			"it didn't work",
		}
	}
	return "OK", nil
}

// ErrorDemo demo
func ErrorDemo(i int) {
	msg, err := run(i)
	if err != nil {
		fmt.Println(msg, err)
	} else {
		fmt.Println(msg)
	}
}
