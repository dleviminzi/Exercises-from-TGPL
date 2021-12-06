package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
)

func main() {
	var input1 = flag.String("i1", "input1", "Problem with first input")
	var input2 = flag.String("i2", "input2", "Problem with second input")
	flag.Parse()

	// calculate the sha of each input
	sha1 := sha256.Sum256([]byte(*input1))
	sha2 := sha256.Sum256([]byte(*input2))

	res := 0
	for i, num1 := range sha1 {
		num2 := sha2[i]
		res += commonBits(num1, num2)
	}

	fmt.Println(res)
}

func commonBits(a, b byte) int {
	res := 0
	// check for mismatch and shift
	for a != 0 {
		if (a&1 == 1 && b&1 == 0) || (a&1 == 0 && b&1 == 1) {
			res += 1
		}
		a = a >> 1
		b = b >> 1
	}
	return res
}
