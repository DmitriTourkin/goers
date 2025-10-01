package main

import "fmt"

var (
	a [2]string
	primes = [...]int{2, 3, 5, 7, 11, 13}
)

func ArrayShow() {
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	var s []int = primes[1:4]
	fmt.Println(s)
}
