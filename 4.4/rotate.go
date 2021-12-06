package main

import "fmt"

// a = array to rotate, p = number of positions
func rotate(a []int, p int) {
	wrapSection := len(a) - p
	holder := []int{}
	j := 0

	for i := 0; i < len(a); i++ {
		if i >= wrapSection {
			a[i] = holder[j]
			j += 1
		} else if i < p {
			holder = append(holder, a[i])
			a[i] = a[i+p]
		} else {
			a[i] = a[i+p]
		}
	}
}

func rotate2(a []int, p int) []int {
	return append(a[p:], a[:p]...)
}

func main() {
	l := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%p", l) // address will remain the same
	rotate(l, 2)
	fmt.Printf("%p", l) // same address
	fmt.Println(l)

	l2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%p", l2) // address will change
	l2 = rotate2(l2, 2)
	fmt.Printf("%p", l2) // new address
	fmt.Println(l2)
}

//
// [1, 2, 3, 4, 5, 6] and p=2
// we want to take the first two elements 1, 2 and throw them to the back
// thus we know that the spots of 5, 6 will be overwritten with non-adjacent data
//
