package main

import "fmt"

// The data values generated in the range clause are copies of the actual collection elements,
// these does not reference original item. Which means that updating the values
// will not change the original data. You should make use of
// index operator on collection to update original value.
func main() {
	trap()
	alternative()
}

func trap() {
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		v *= 5 //original item is not changed
	}

	fmt.Println(s) // prints [1 2 3 4 5]
}

func alternative() {
	s := []int{1, 2, 3, 4, 5}
	for i := range s {
		s[i] *= 5
	}

	fmt.Println(s) // prints data: [5 10 15 20 25]
}

// Output of program:
// [1 2 3 4 5]
// [5 10 15 20 25]
