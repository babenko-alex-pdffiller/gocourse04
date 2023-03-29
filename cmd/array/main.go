package main

import (
	"fmt"
)

func main() {
	appleBoxes := [10]int{} // not a pointer
	fmt.Println(appleBoxes)
	fmt.Println(len(appleBoxes))
	fmt.Println(cap(appleBoxes))

	bananaBoxes := [10]int{}
	bananaBoxes[1] = 100
	bananaBoxes[3] = 200
	fmt.Println(bananaBoxes)
	fmt.Println(bananaBoxes[0], bananaBoxes[1], bananaBoxes[2], bananaBoxes[3])

	potatoBoxes := [10]int{1, 1, 4, 1, 0, 1, 7, 1, 1, 1}
	tomatoBoxes := [...]int{1, 2, 3}
	fmt.Println(potatoBoxes, tomatoBoxes)
}
