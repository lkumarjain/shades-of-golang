package main

import "fmt"

// Map returns a zero-value of a type for non-existing keys.
// Checking for the appropriate zero value can be used to determine,
// if the map record exists or not, but it's not always reliable
// (e.g., what if you have a map of Boolean where the zero value is false).
// The most reliable way to know, if a given map record exists or not
// is to check the second value returned by the map access operation.
func main() {
	trap()
	alternative()
}

func trap() {
	m := map[int]bool{1: true, 2: false, 3: false}

	if v := m[4]; v { // Zero-value of v is false
		fmt.Println("Trap ", "no entry")
	} else {
		fmt.Println("Trap ", "entry exist") // Prints: Trap  entry exist
	}
}

func alternative() {
	m := map[int]bool{1: true, 2: false, 3: false}

	if _, ok := m[4]; !ok { // value of Ok is false
		fmt.Println("Alternative ", "no entry") // Prints: Alternative  no entry
	} else {
		fmt.Println("Alternative ", "entry exist")
	}
}

// Output :
// Trap  entry exist
// Alternative  no entry
