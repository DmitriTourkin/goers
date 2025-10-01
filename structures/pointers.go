package main

import "fmt"

// Pointer — holds memory address of a value
// *T — pointer to a T value. Zero value = nil
// & operator — generates pointer to its operand

func PointersAndValues() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p) // read i through the pointer
	*p = 21 // set i through the pointer
	fmt.Println(i)

	p = &j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)
}