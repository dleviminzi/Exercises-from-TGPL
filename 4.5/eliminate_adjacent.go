package main

import "fmt"

//
// We will go through until we find a duplicate.
// Then we will remove the duplicate.
// Slice at the index where we removed it.
// Then recurse down.
//

func eliminate(s []string) []string {
	for i, word := range s {
		if i+1 < len(s) && s[i+1] == word {
			// loop until repeats are finished
			var j int
			for j = i + 1; j < len(s); j++ {
				if word != s[j] {
					break
				}
			}
			// append first of repeats and skip the rest
			s = append(s[:i+1], s[j:]...)
		}
	}
	return s
}

func main() {
	test := []string{"steve", "daniel", "daniel", "daniel", "jon", "brad", "landon", "landon"}
	fmt.Println(eliminate(test))
}
