package main

import "fmt"

func main() {
	var arrayInt [3]int

	arrayInt[0] = 11
	arrayInt[1] = 22

	for i, element := range arrayInt {
		fmt.Printf("arrayInt[%v]=%v\n", i, element)
	}
}
