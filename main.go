package main

import "fmt"

const (
	// We need another constant..obviously.
	urgentconstant1 = 1

	// Create constant using 'iota' and start from 1 (iota starts from zero)
	constant1 = iota + 1

	// Create another. It will naturally consume from iota.
	constant2
)

func main() {
	// Things check out as expected
	fmt.Println("Constant 1: ", constant1)
	fmt.Println("Constant 2: ", constant2)

	// Needed this for some reason
	fmt.Println("Urgent Constant: ", urgentconstant1)

	// Output
	// -------
	// Constant 1:  2
	// Constant 2:  3
	// Urgent Constant:  1
	//
}
