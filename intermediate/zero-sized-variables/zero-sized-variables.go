package main

import "fmt"

// Shouldn’t two different variables have different addresses?
// Well, it's not the case with Go 😊,
// if you have zero-sized variables they might share the exact same address in memory.
func main() {
	learning()
}

type data struct{}

func learning() {
	x := &data{}
	y := &data{}

	fmt.Printf("Same address - x=%p y=%p\n", x, y)
	//prints: same address - a=0x586450 b=0x586450
}

// Output:
// Same address - x=0x586450 y=0x586450
