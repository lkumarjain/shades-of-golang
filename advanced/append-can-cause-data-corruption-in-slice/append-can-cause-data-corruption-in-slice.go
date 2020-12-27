package main

import "fmt"

// Slices are backed by an array and capacity as the original.
// So, if you change an element in the slice, the original contents are modified as well.
// At some time adding data to one of the slices may result in a to a new array allocation,
// if the original array can't hold more data as a result other slices
// will point to the old array (with old data) whereas this can point to new instance of slice.
func main() {
	trap()
	alternative()
}

func trap() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s1), cap(s1), s1) // prints 5 5 [1 2 3 4 5]

	s2 := s1[1:]                      // Backed same array as the original
	fmt.Println(len(s2), cap(s2), s2) // prints 4 4 [2 3 4 5]

	for i := range s2 {
		s2[i] += 10 // Modifies slice s1 along with s2
	}

	// Still referencing the same array as a result updation happen on both
	fmt.Println(len(s1), cap(s1), s1) // prints 5 5 [1 12 13 14 15]
	fmt.Println(len(s2), cap(s2), s2) // prints 4 4 [12 13 14 15]

	s2 = append(s2, 16)

	for i := range s2 {
		s2[i] += 40
	}

	// s1 is now "stale" as data corruption happened
	fmt.Println(len(s1), cap(s1), s1) // prints 5 5 [1 12 13 14 15]
	fmt.Println(len(s2), cap(s2), s2) // prints 5 8 [52 53 54 55 56]
}

func alternative() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s1), cap(s1), s1) // prints 5 5 [1 2 3 4 5]

	s2 := make([]int, 4)              // Creating a new slice
	copy(s2, s1[1:])                  // Coping s1 data into s2 to avoid any collision
	fmt.Println(len(s2), cap(s2), s2) // prints 4 4 [2 3 4 5]

	for i := range s2 {
		s2[i] += 10 // Modifies s2 slice only; s1 is intect
	}

	// Both are referancing to different array as a result not impact on each other
	fmt.Println(len(s1), cap(s1), s1) // prints 5 5 [1 2 3 4 5]
	fmt.Println(len(s2), cap(s2), s2) // prints 4 4 [12 13 14 15]
}

// Output :
// 5 5 [1 2 3 4 5]
// 4 4 [2 3 4 5]
// 5 5 [1 12 13 14 15]
// 4 4 [12 13 14 15]
// 5 5 [1 12 13 14 15]
// 5 8 [52 53 54 55 56]
// 5 5 [1 2 3 4 5]
// 4 4 [2 3 4 5]
// 5 5 [1 2 3 4 5]
// 4 4 [12 13 14 15]
