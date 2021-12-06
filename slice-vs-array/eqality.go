package main

import "fmt"

func main() {
	// These are arrays
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 4, 5}

	// These are slices
	aS := []int{1, 2, 3, 4, 5}
	bS := []int{1, 2, 3, 4, 5}

	fmt.Println(&a == &b)
	fmt.Println(a == b)
	// wont work -> fmt.Println(aS == bS)

	fmt.Println(&aS == &bS)
}
