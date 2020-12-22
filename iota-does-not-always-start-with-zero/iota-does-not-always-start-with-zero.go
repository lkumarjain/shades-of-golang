package main

import "fmt"

// At a first glance, It seems that iota identifier is an increment operator
// starts form zero and continues to increase from there,
// well this is not always true in Go.
// The value of iota is identified by an index operator
// for the current line in the constant declaration block,
// which means if the use of iota is not the first line in
// the constant declaration block the initial value may not be zero.
const (
	a = iota // Iota start with 0
	b = iota // Iota assigns next value 1
)

const (
	ignored = "processing" // Constant Index 0
	x       = iota         // Iota starts with next constant Index 1
	y       = iota         // Iota assigns next value 2
)

func main() {
	fmt.Println(a, b) //prints: 0 1
	fmt.Println(x, y) //prints: 1 2
}

// Output or the program is:
// 0 1
// 1 2
