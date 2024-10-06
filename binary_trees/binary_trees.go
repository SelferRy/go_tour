package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

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

func testWalk() {
	var t = tree.New(3)
	var ch = make(chan int, 10)
	go Walk(t, ch)
	fmt.Println(<-ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	n := 10
	ch1 := make(chan int, n)
	ch2 := make(chan int, n)
	defer close(ch1)
	defer close(ch2)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < n; i++ {
		val1 := <-ch1
		val2 := <-ch2
		if val1 != val2 {
			return false
		}
	}
	return true
}

func testSame() {
	var t1, t2 *tree.Tree = tree.New(1), tree.New(1)
	var result bool = Same(t1, t2)
	fmt.Println(result)
}

func main() {
	testWalk()
	testSame()
}
