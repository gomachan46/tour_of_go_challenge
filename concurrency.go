package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// 1.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t != nil {
			walker(t.Left)
			ch <- t.Value
			walker(t.Right)
		}
	}

	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// 3.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	// 2. テストとは、といった感じだけど
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	// 4.
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}