package main

import "fmt"

// Another most common gotcha for the developer accustomed to a foreach loop in Java.
// In Go, the range clause is different, it generates two values:
// first value is the index and the second one is data.
func main() {
	trap()
	alternative()
}

func trap() {
	x := []string{"a", "b", "c"}

	for v := range x { // Running range loop and capturing index
		fmt.Println(v) //prints 0, 1, 2
	}
}

func alternative() {
	x := []string{"a", "b", "c"}

	for _, v := range x { // Ignoring index and capturing value
		fmt.Println(v) //prints a, b, c
	}
}
