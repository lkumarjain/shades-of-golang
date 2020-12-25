package main

import "fmt"

// A nil map behaves like an empty map when reading,
// but attempts to write to a nil map will cause a runtime panic;
// don't do that. Always initialize a map, using the built-in `make` function.
func main() {
	trap()
	alternative()
}

func trap() {
	var m map[string]int // Initialization is not done
	i, ok := m["one"]    // Reading on nil map does not panic
	if !ok {
		m["one"] = 1 // runtime panic as trying to write on nil map
	}
	fmt.Println(i)
}

func alternative() {
	var m map[string]int = make(map[string]int) // Initializing to avoid panic
	m["one"] = 1                                // Write operation is success
}

// Output of a program
// panic: assignment to entry in nil map
