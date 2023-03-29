package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	// 1
	northExpress := make( // pointer
		[]int,
		10, // len
		12, // cap
	)
	fmt.Println(northExpress)
	fmt.Println(len(northExpress))
	fmt.Println(cap(northExpress))

	// 2
	southExpress := []int{1}
	southExpressMirror := southExpress[0:1]
	fmt.Printf("%p\n", &southExpress)
	fmt.Printf("%p\n", &southExpressMirror)

	southExpress = []int{1}
	southExpressMirror = southExpress[0:1]
	southExpressMirror[0] = 0

	fmt.Println(southExpress)

	southExpress[0] = 1

	fmt.Println(southExpressMirror)

	// 3
	westExpress := make([]int, 1, 2)
	westExpress[0] = 1
	westExpressPart := westExpress[0:1]

	westExpress = append(westExpress, 2)
	westExpress[0] = 10

	fmt.Println(westExpress, westExpressPart) // [10 2] [1]

	// 4
	eastExpress := []int{1, 2, 3, 4, 5, 6, 7}
	eastExpressPart := eastExpress[3:5]

	fmt.Println(eastExpress, cap(eastExpress))
	fmt.Println(eastExpressPart, cap(eastExpressPart))

	eastExpressPart[0] = 1_000_000
	fmt.Println(eastExpressPart)

	// 5
	almostLastExpress := []int{1, 2, 3}
	almostLastExpressCopy := make([]int, 3, 3)
	copy(almostLastExpressCopy, almostLastExpress)

	almostLastExpress[2] = 2
	almostLastExpressCopy[0] = 0
	fmt.Println(almostLastExpress, almostLastExpressCopy)

	// 7
	lastExpress := []int{1, 2, 3, 4, 5, 6, 7}
	lastExpressPart := lastExpress[0:1:5]
	fmt.Println(len(lastExpress), cap(lastExpress))
	fmt.Println(len(lastExpressPart), cap(lastExpressPart))
}
