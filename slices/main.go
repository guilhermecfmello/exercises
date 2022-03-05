package main

import (
	"fmt"
	"reflect"
)

func main() {

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	// stringPointer = &names
	// stringPointer[0] = "Guilherme"

	// printVariableAndType(names)

	// a := names[0:2]
	b := names[1:3]

	// printVariableAndType(a)
	// printVariableAndType(b)
	b[0] = "XXX"
	// printVariableAndType(a)
	// printVariableAndType(b)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	slice := []int{1, 2, 3, 4, 5}
	printSlice(slice)
	slice = make([]string, 10, 20)
	fmt.Print(slice)
	// printSlice(slice)
}

func printVariableAndType(x interface{}) {
	fmt.Println("Type: ", reflect.TypeOf(x))
	fmt.Println(x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
