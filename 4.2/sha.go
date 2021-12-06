package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	shaNum := flag.String("num", "256", "Provide specific Sha number you want")
	flag.Parse()
	if !ValidSha(*shaNum) {
		fmt.Println("Only 256, 384, and 512 are options")
		os.Exit(1)
	}

	if len(os.Args) != 3 {
		fmt.Println("Must provide num and data")
		os.Exit(1)
	}
	d := os.Args[2]

	if *shaNum == "256" {
		h := sha256.New()
		h.Write([]byte(d))
		hash := hex.EncodeToString(h.Sum(nil))
		fmt.Println(d, hash)
	} else if *shaNum == "384" {
		h := sha512.New384()
		h.Write([]byte(d))
		hash := hex.EncodeToString(h.Sum(nil))
		fmt.Println(d, hash)
	} else {
		h := sha512.New()
		h.Write([]byte(d))
		hash := hex.EncodeToString(h.Sum([]byte(d)))
		fmt.Println(d, hash)
	}
}

func ValidSha(s string) bool {
	switch s {
	case
		"256",
		"384",
		"512":
		return true
	}
	return false
}
