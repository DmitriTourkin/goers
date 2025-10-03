package main

import "fmt"

var powSlice = []int{1, 2, 4, 8, 16, 32, 64, 128}

func RangingExample() {
	for i, v := range powSlice {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}