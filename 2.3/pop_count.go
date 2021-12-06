package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/* bit lost on this one */
func PopCountLoop(x uint64) int {
	res := 0
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		res += int(pc[i])
	}
	return res
}

func main() {
	var x uint64
	x = 88

	fmt.Println(PopCount(x))
	fmt.Println(PopCountLoop(x))
}
