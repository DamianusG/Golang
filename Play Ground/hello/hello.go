package main

import (
	"fmt"
)

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 4000000000.
	fmt.Printf("%.2fKB\n", fileSize/KB)
	fmt.Printf("%v", KB)
	//s := "this is a string"
	//s[2] = strconv.Atoi("u")
	// fmt.Printf("%v, %T\n", s[2], s[2])
	// fmt.Printf("%v, %T\n", s[2], s[2])
	// fmt.printf("%v, %T\n", s, s)
}
