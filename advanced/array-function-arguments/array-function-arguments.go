package main

import "fmt"

// As a Java developer, you may think, when you pass arrays to functions
// the functions reference the same memory location,
// so they can update the original data. But it's different in Go,
// arrays are values and does not share memory location,
// so when you pass arrays to functions.
// They get their own copy of the original array data.
// This can be a problem, if you are trying to update the array data.
func main() {
	trap()
	alternative()
	viaSlice()
}

func trap() {
	a := [6]int{1, 2, 3, 4, 5, 6}

	func(arr [6]int) {
		arr[0] = 7                // Updating a copy of array
		fmt.Println("Trap ", arr) // prints [7 2 3 4 5 6]
	}(a) // Passing array value to function

	fmt.Println("Trap ", a) // prints [1 2 3 4 5 6] (not ok if you need [7 2 3 4 5 6])
}

func alternative() {
	a := [6]int{1, 2, 3, 4, 5, 6}

	func(arr *[6]int) {
		(*arr)[0] = 7                    // Updating a copy of array
		fmt.Println("Alternative ", arr) // prints &[7 2 3 4 5 6]
	}(&a) // Passing array address to function

	fmt.Println("Alternative ", a) // prints [7 2 3 4 5 6]
}

func viaSlice() {
	a := []int{1, 2, 3, 4, 5, 6}

	func(arr []int) {
		arr[0] = 7                 // Updating a copy of slice
		fmt.Println("Slice ", arr) // prints [7 2 3 4 5 6]
	}(a) // Passing slice address to function

	fmt.Println("Slice ", a) // prints [7 2 3 4 5 6]
}

// Output :
// Trap  [7 2 3 4 5 6]
// Trap  [1 2 3 4 5 6]
// Alternative  &[7 2 3 4 5 6]
// Alternative  [7 2 3 4 5 6]
// Slice  [7 2 3 4 5 6]
// Slice  [7 2 3 4 5 6]
