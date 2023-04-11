package main

import "fmt"

func main() {
	tree := []int{1}

	for i := 1; i <= 5; i++ {
		prevLen := len(tree)
		newLen := prevLen + 2*prevLen
		branchSizes := make([]int, newLen)

		for j := 0; j < prevLen; j++ {
			branchSizes[2*j] = tree[j]
			branchSizes[2*j+1] = tree[j]
		}

		tree = branchSizes
	}

	fmt.Println(tree)
}
