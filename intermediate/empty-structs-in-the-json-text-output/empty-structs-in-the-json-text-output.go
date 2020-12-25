package main

import (
	"encoding/json"
	"fmt"
)

// Person - Struct to hold Person information
type Person struct {
	Name string
	age  int
}

// Only exported fields of a Go struct will be present in the JSON output.
// You can specify the JSON field name explicitly with a `json: tag`.
func main() {
	p := Person{"Person Name", 22}
	fmt.Printf("%#v\n", p) //prints: main.Person{Name:"Person Name", age:22}

	encoded, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(string(encoded)) //prints: {"Name":"Person Name"}

	var o Person
	json.Unmarshal(encoded, &o)

	fmt.Printf("%#v\n", o) //prints: main.Person{Name:"Person Name", age:0}
}

// Output of program:
// main.Person{Name:"Person Name", age:22}
// {"Name":"Person Name"}
// main.Person{Name:"Person Name", age:0}
