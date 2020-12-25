package main

import "fmt"

type A struct {
	Name string
}

func (a *A) Go() {
	fmt.Printf("GO: %v\n", a)
}

func Test(fn func()) {
	fn()
}

// Go developers many times faces the issue of dereferencing of a `nil` pointer.
// An uninitialized pointer is nil, and you can’t easily follow the nil pointer.
// If `x is nil`, an attempt to evaluate `*x` will cause a run-time panic.
func main() {
	aa := &A{Name: "FOO"}
	bb := (*A)(nil)
	cc := &A{}
	Test(aa.Go)

	// The variable bb is nil so calling bb.Go() method is expected to cause a panic
	// "runtime error: invalid memory address or nil pointer dereference",
	// but the method call succeeds.
	// detailed explaination available in
	// https://lkumarjain.blogspot.com/2020/01/why-calling-method-on-nil-struct.html
	Test(bb.Go)

	Test(cc.Go)
}
