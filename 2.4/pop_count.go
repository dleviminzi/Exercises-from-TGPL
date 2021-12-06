package main

import "fmt"

func main() {
	// Another loop based pop count
	count := 0
	y := 88
	for y != 0 {
		count += int(y & 1)
		y = y >> 1
	}
	fmt.Println(count)
}
