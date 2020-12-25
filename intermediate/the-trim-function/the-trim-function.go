package main

import (
	"fmt"
	"strings"
)

// The Trim, Trim Left, and Trim Right functions strip
// all Unicode code points contained in a cut-set.
// To strip a trailing string, you should make use of Trim Suffix or Trim Prefix.
func main() {
	fmt.Println(strings.TrimRight("CABBBBBBA", "BA")) // Output: ""
	fmt.Println(strings.TrimSuffix("CABBA", "BA"))    // Output: "AB"
}

// Output is:
// C
// CAB
