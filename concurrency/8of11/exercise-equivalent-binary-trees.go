package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//Walk the tree
func DoWalk(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go DoWalk(t1, ch1)
	ch2 := make(chan int)
	go DoWalk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 && !ok2 || !ok1 && ok2 {
			return false
		}

		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)
	fmt.Printf("Tree 1: %v\n", t1)
	fmt.Printf("Tree 2: %v\n", t2)
	fmt.Printf("Tree 3: %v\n", t3)
	fmt.Printf("Do the trees t1 and t2 store the same values? -> %v\n", Same(t1, t2))
	fmt.Printf("Do the trees t1 and t3 store the same values? -> %v\n", Same(t1, t3))
}
