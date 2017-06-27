package main

import "fmt"

const (

	// Create constants using 'iota' and start from 1 (iota starts from zero)
	GET = iota + 1
	SET
	DEL
)

func main() {
	// Things check out as expected
	fmt.Println("GET: ", GET)
	fmt.Println("SET: ", SET)
	fmt.Println("DEL: ", DEL)

	// Output
	// -------
	// GET:  1
	// SET:  2
	// DEL:  3
	//
}
