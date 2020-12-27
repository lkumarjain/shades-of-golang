package main

import (
	"fmt"
	"time"
)

// This is yet another most common pitfall in Go,
// the iteration variables in for statements are reused across iterations.
// Which means that each closure created in the for loop will reference the same variable.
// You can solve this by creating a local variable in the loop or
// other solution could be to pass variable as a parameter in the closure function.
// We can detect this early in the program by using [Go Race Detector](https://blog.golang.org/race-detector).

type field struct {
	name string
}

func (p *field) print() {
	fmt.Print(" ", p.name, " ")
}

func main() {
	trap()
	alternative1()
	alternative2()
}

func trap() {
	d := []field{{"N1"}, {"N2"}, {"N3"}, {"N4"}, {"N5"}}
	fmt.Print("Trap")

	for _, v := range d {
		go v.print()
	}

	time.Sleep(3 * time.Second) //prints: Trap N5  N5  N5  N5  N5
	fmt.Println()
}

func alternative1() {
	d := []field{{"N1"}, {"N2"}, {"N3"}, {"N4"}, {"N5"}}
	fmt.Print("Alternative1")

	for _, v := range d {
		v := v //Create a shadow copy
		go v.print()
	}

	time.Sleep(3 * time.Second) //prints: Alternative1 N5  N2  N1  N4  N3
	fmt.Println()
}

func alternative2() {
	d := []*field{{"N1"}, {"N2"}, {"N3"}, {"N4"}, {"N5"}}
	fmt.Print("Alternative2")

	for _, v := range d {
		go v.print()
	}

	time.Sleep(3 * time.Second) //prints: Alternative2 N5  N3  N4  N2  N1
	fmt.Println()
}

// Output :
// Trap N5  N5  N5  N5  N5
// Alternative1 N5  N2  N1  N4  N3
// Alternative2 N5  N3  N4  N2  N1
