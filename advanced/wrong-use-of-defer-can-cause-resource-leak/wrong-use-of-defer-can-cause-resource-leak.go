package main

import (
	"fmt"
	"net/http"
)

// You might use a defer in the code block and think that this will be called at the end of function to clean resources.
// This can eventually cause a resource leak if you are running a long for loop and calling
// defer in the code block to release resources. This can be avoided,
// by warping code block to function block or use direct clean-up instead of defer statement.
func main() {
	trap()
	alternative()
}

var targets = []string{
	"https://lkumarjain.blogspot.com",
	"https://github.com/lkumarjain",
	"https://lkumarjain.blogspot.com/2020/11/golang-from-java-developer-point-of-view.html",
}

func trap() {
	for _, target := range targets {
		r, err := http.Get(target)
		fmt.Println("Trap: Requesting target:", target)
		if err != nil {
			// This can eventually print: error: too many open files
			fmt.Println("Trap: bad target:", target, "error:", err)
			break
		}
		defer func(resp *http.Response, t string) {
			fmt.Println("Trap: Closing Body target:", t)
			resp.Body.Close()
		}(r, target) // will not be closed at the end of this code block
	}
}

func alternative() {
	for _, target := range targets {
		func() {
			r, err := http.Get(target)
			fmt.Println("Alternative: Requesting target:", target)
			if err != nil {
				fmt.Println("Alternative: bad target:", target, "error:", err)
			}
			defer func() {
				fmt.Println("Alternative: Closing Body target:", target)
				r.Body.Close()
			}() // will Be called for each funciton call
		}()
	}
}

// Output of this program:
// Trap: Requesting target: https://lkumarjain.blogspot.com
// Trap: Requesting target: https://github.com/lkumarjain
// Trap: Requesting target: https://lkumarjain.blogspot.com/2020/11/golang-from-java-developer-point-of-view.html
// Trap: Closing Body target: https://lkumarjain.blogspot.com/2020/11/golang-from-java-developer-point-of-view.html
// Trap: Closing Body target: https://github.com/lkumarjain
// Trap: Closing Body target: https://lkumarjain.blogspot.com
// Alternative: Requesting target: https://lkumarjain.blogspot.com
// Alternative: Closing Body target: https://lkumarjain.blogspot.com
// Alternative: Requesting target: https://github.com/lkumarjain
// Alternative: Closing Body target: https://github.com/lkumarjain
// Alternative: Requesting target: https://lkumarjain.blogspot.com/2020/11/golang-from-java-developer-point-of-view.html
// Alternative: Closing Body target: https://lkumarjain.blogspot.com/2020/11/golang-from-java-developer-point-of-view.html
