package main

import (
	"fmt"
	"math"
)

func loops() {
	fmt.Println("Loops Function")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func whileLoop() {
	fmt.Printf("For/While loop Example")
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	fmt.Println("Sqrt Function")
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func practiseLoop() {
	fmt.Println("Practice Loops Function")
	var state bool = true

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i % 2 == 0 {
			fmt.Println(state)
		}
	}
}

func AllFns() {
	whileLoop()
	loops()
	practiseLoop()
}