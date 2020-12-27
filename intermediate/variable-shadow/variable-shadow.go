package main

import "fmt"

// This is a very common trap even for experienced Go developers.
// It's easy to make and it could be hard to spot.
// You can make use of linters `vet` or
// [go-nyet](https://github.com/barakmich/go-nyet) `for more aggressive checks`
// to detect such errors in the program.
func main() {
	trap()
	alternative()
}

func trap() {
	v := "Trap 1"
	fmt.Println(v) // prints "Trap 1"
	{
		fmt.Println(v) // prints "Trap 1"
		v := "Trap 2"  // Variable shadow
		fmt.Println(v) // prints "Trap 2"
	}
	fmt.Println(v) // prints "Trap 1" (bad if you need 2)
}

func alternative() {
	v := "Alternative 1"
	fmt.Println(v) // prints "Alternative 1"
	{
		fmt.Println(v)      // prints Alternative 1
		v = "Alternative 2" // Removed Variable shadow
		fmt.Println(v)      //prints Alternative 2
	}
	fmt.Println(v) // prints Alternative 2
}

// Output:
// Trap 1
// Trap 1
// Trap 2
// Trap 1
// Alternative 1
// Alternative 1
// Alternative 2
// Alternative 2
