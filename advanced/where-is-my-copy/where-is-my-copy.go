package main

import "fmt"

var (
	source = []int{1, 2, 3, 4, 5}
)

// The copy function in the Go, copies minimum elements of the source to destination,
// to ensure correct coping, you should allocate sufficient destination slice.
// The copy number of elements copied. Alternatively,
// you can use append function in Go to copy array elements,
// make a note that, size (capacity) of the append slice will
// be larger than the length of slice.
func main() {
	trap()         // Prints: Trap: Source:  [1 2 3 4 5]  Destination:  []
	alternative1() // Prints: Alternative1: Source: [1 2 3 4 5] Destination:  [1 2 3 4 5]  (Copied 5 numbers)
	alternative2() // Prints:  Alternative2: Source: [1 2 3 4 5] Destination:  [1 2 3 4 5]
}

func trap() {
	var destination []int
	copy(destination, source) // Copy elements to destination from source.
	fmt.Println("Trap: Source: ", source, " Destination: ", destination)
}

func alternative1() {
	var destination []int
	destination = make([]int, len(source))
	n := copy(destination, source)
	fmt.Println("Alternative1: Source:", source, "Destination: ", destination, " (Copied", n, "numbers)")
}

func alternative2() {
	var destination []int
	destination = append(destination, source...)
	fmt.Println("Alternative2: Source:", source, "Destination: ", destination)
}

// Output of this program:
// Trap: Source:  [1 2 3 4 5]  Destination:  []
// Alternative1: Source: [1 2 3 4 5] Destination:  [1 2 3 4 5]  (Copied 5 numbers)
// Alternative2: Source: [1 2 3 4 5] Destination:  [1 2 3 4 5]
