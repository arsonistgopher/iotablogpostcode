package main

import "fmt"

// This block is just for the JSON-API
const (
	GET = iota + 1
	SET
	DEL
)

// This block is for other consts not relying on iota
const (
	ALARM_LEVEL = 3
)

func main() {
	// Things check out as expected
	fmt.Println("GET: ", GET)
	fmt.Println("SET: ", SET)
	fmt.Println("DEL: ", DEL)
	fmt.Println("ALARM_LEVEL: ", ALARM_LEVEL)

	// Output
	// -------
	// GET:  1
	// SET:  2
	// DEL:  3
	// ALARM_LEVEL:  3
	//
}
