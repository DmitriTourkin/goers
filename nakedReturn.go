package main

import "fmt"

func splitNakedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func NakedResultExample() {
	fmt.Println(splitNakedReturn(16))
}