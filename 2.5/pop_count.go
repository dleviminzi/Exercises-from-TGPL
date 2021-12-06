package main

import "fmt"

func main() {
	y := 88
	count := 0

	// Loop based pop count
	for y != 0 {
		y = y & (y - 1)
		count += 1
	}
	fmt.Println(count)
}
