package main

import "fmt"

// The new slice created out of an existing slice will reference the array of the original slice.
// This behavior can lead to unexpected memory usage,
// if your application allocates large slices and creates new slices
// from them to refer to small sections of the original data.
// This can be avoided by making a copy of the data needed from the temporary slice.
func main() {
	trap()
	alterrnative()
}

func trap() {
	raw := make([]byte, 100)
	fmt.Println("Trap", len(raw), cap(raw), &raw[0]) // prints: Trap 100 100 0xc00001a150
	data := raw[:10]
	fmt.Println("Trap", len(data), cap(data), &data[0]) // prints: Trap 10 100 0xc00001a150
}

func alterrnative() {
	raw := make([]byte, 100)
	fmt.Println("Alterrnative", len(raw), cap(raw), &raw[0]) // prints: Alterrnative 100 100 0xc00010c070
	data := make([]byte, 10)
	copy(data, raw[:10])
	fmt.Println("Alterrnative", len(data), cap(data), &data[0]) // prints: Alterrnative 10 10 0xc000112008
}

// Output of a program
// Trap 100 100 0xc00010c000
// Trap 10 100 0xc00010c000
// Alterrnative 100 100 0xc00010c070
// Alterrnative 10 10 0xc000112008
