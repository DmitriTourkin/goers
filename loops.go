package main

import (
	"fmt"
	"math"
)

func loops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func whileLoop() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func allFns() {
	whileLoop()
	loops()
	sqrt()
}