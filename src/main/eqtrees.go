package main

import "fmt"

// Tree export
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traversal(tree *Tree, ch chan int) {
	if tree != nil {
		traversal(tree.Left, ch)
		ch <- tree.Value
		traversal(tree.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(tree *Tree, ch chan int) {
	traversal(tree, ch)
	close(ch)
}

// WalkDemo demo
func WalkDemo() {
	treech := make(chan int)
	t := Tree{
		&Tree{Value: 1},
		2,
		&Tree{Value: 3}}
	t2 := Tree{
		&Tree{Value: 1},
		42,
		&Tree{Value: 3}}
	t3 := Tree{
		Left:  &Tree{Value: 1},
		Value: 2}
	go Walk(&t, treech)
	for i := range treech {
		fmt.Println(i)
	}
	fmt.Println("Same:", Same(&t, &t))
	fmt.Println("Same:", Same(&t, &t2))
	fmt.Println("Same:", Same(&t, &t3))
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		e1, ok1 := <-ch1
		e2, ok2 := <-ch2
		if ok1 && ok2 {
			if e1 != e2 {
				return false
			}
		} else if !ok1 && !ok2 { // both closed
			return true
		} else {
			return false
		}
	}
}
