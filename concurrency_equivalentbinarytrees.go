package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	Walk_Helper(t, ch) //Functionality here so that we can close the channel afterwards
	close(ch)
}

func Walk_Helper(t *tree.Tree, ch chan int){
	if (t.Left != nil) {
		Walk_Helper(t.Left, ch)
	}
	
	ch <- t.Value
	
	if (t.Right != nil) {
		Walk_Helper(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	var arr1 []int
	var arr2 []int
	
	for i := range ch1 {
		arr1 = append(arr1, i)
	}
	
	for i := range ch2 {
		arr2 = append(arr2, i)
	}

	if (len(arr1) != len(arr2)) {
		return false
	}
	
	for i := 0; i < len(arr1); i ++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	
	return true
	
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
