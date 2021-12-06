package main

import "fmt"

// reverse reverses a slice of ints in place.
// given by book
func reverseTGPL(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//
// Goal of this exercise is to rewrite the above function using an array
// pointer instead of using a slice.
//
func reverse(s *[]int) {
	fmt.Println(len(*s))
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	// Technically this is a slice that we are defining here...
	ex1 := []int{1, 2, 3, 4, 5}
	reverse(&ex1)
	fmt.Println(ex1)
}

// I think that part of the point of the exercise is to make you realize that the passing of the array pointer is
// either restrictive or messy. Either we need to know the size of the array whose pointer is being sent/received
// or we need to use slices anyway.
