package main

import "fmt"

func main() {
	var array [2]int

	// zero value of an array is zero-valued elements
	fmt.Printf("array       : %#v\n", array)

	// nope: arrays are fixed length
	// array[2] = 0

	var slice []int

	// zero value of a slice is nil
	fmt.Println("slice == nil?", slice == nil)

	// nope: they don't exist:
	// _ = slice[0]
	// _ = slice[1]

	// len function still works though
	fmt.Println("len(slice)  :", len(slice))

	// array's length is part of its type
	fmt.Printf("array's type: %T\n", array)

	// whereas, slice's length isn't part of its type
	fmt.Printf("slice's type: %T\n", slice)
}
