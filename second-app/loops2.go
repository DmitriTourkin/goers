package main

import (
	"fmt"
	"math"
)

func Pow(x, n, lim float64) float64 {
	fmt.Println("Pow Function")
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}