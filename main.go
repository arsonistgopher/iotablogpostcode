package main

import "fmt"

const (
	// Const block for JSON-API
	GET = iota + 1
	SET
	DEL
)

const (
	// Other consts
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
	// GET:  2
	// SET:  3
	// DEL:  4
	// ALARM_LEVEL:  3
	//
}
