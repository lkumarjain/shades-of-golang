package main

import "fmt"

// You might think we need to add break statement in the switch to move out,
// well this is quite different in Go, we don’t need to add break in the case block,
// as a result fall-through does not work as expected.
// You can force this in the case blocks to fall through by using the fall-through statement
// at the end of each case block or rewrite switch statement to use expression lists in the case blocks.
func main() {
	trap()
	alternative1()
	alternative2()
}

func trap() {
	isValidNumber := func(i int) bool {
		switch i {
		case 1: // Ignored as fallthrough is not done
		case 2:
			return true
		}
		return false
	}

	fmt.Println("Trap", isValidNumber(1)) //prints: Trap true (ok)
	fmt.Println("Trap", isValidNumber(2)) //prints: Trap false (not ok)
	fmt.Println("Trap", isValidNumber(3)) //prints: Trap false (ok)
}

func alternative1() {
	isValidNumber := func(i int) bool {
		switch i {
		case 1, 2: // rewritten case statement
			return true
		}
		return false
	}

	fmt.Println("Alternative1", isValidNumber(1)) //prints: Alternative1 true (ok)
	fmt.Println("Alternative1", isValidNumber(2)) //prints: Alternative1 true (ok)
	fmt.Println("Alternative1", isValidNumber(3)) //prints: Alternative1 false (ok)
}

func alternative2() {
	isValidNumber := func(i int) bool {
		switch i {
		case 1:
			fallthrough // fallthrough in case statement
		case 2:
			return true
		}
		return false
	}

	fmt.Println("Alternative2", isValidNumber(1)) //prints: Alternative2 true (ok)
	fmt.Println("Alternative2", isValidNumber(2)) //prints: Alternative2 true (ok)
	fmt.Println("Alternative2", isValidNumber(3)) //prints: Alternative2 false (ok)
}

// Output :
// Trap false
// Trap true
// Trap false
// Alternative1 true
// Alternative1 true
// Alternative1 false
// Alternative2 true
// Alternative2 true
// Alternative2 false
