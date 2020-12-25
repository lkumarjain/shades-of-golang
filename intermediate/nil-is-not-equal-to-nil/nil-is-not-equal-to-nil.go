package main

import "fmt"

// An interface equals to another interface, only if the concrete value and dynamic type are both nil.
// The same applies to nil value. You can think of the interface value nil as typed,
// and nil without type doesn’t equal nil with type.
// If we convert nil to the correct type, the values are indeed equal.
func main() {
	trap()
	alternative()
}

func trap() {
	work := func(arg int) interface{} {
		var result *struct{} = nil

		if arg > 0 {
			result = &struct{}{}
		}

		return result
	}

	// 'res' is not 'nil', but its value is 'nil'
	if res := work(-1); res != nil {
		fmt.Println("Trap", "good result:", res) //Prints: Trap good result: <nil>
	} else {
		fmt.Println("Trap", "bad result (res is nil)")
	}
}

func alternative() {
	work := func(arg int) interface{} {
		var result *struct{} = nil

		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil //return an explicit 'nil'
		}

		return result
	}

	if res := work(-1); res != nil {
		fmt.Println("Alternative", "good result:", res)
	} else {
		fmt.Println("Alternative", "bad result (res is nil)") // Prints: Alternative bad result (res is nil)
	}
}

// Output of program:
// Trap good result: <nil>
// Alternative bad result (res is nil)
