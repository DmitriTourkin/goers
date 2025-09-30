package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool

// Basic Types
var (
	ToBe bool = false
	line string = "Hello!"
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)


func variablesWithInitializers() {
	var i, j int = 1, 2
	var n, m = 2, 3
	fmt.Println(i, j, n, m)
}

func showGolangTypes() {
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v\n", z, z)
}

func insideFnShortAssign() {
	k := 5
	c, python, java := true, false, "no!"
	fmt.Println(k, c, python, java)
}

func main() {
	var i int
	fmt.Println(i, c, python, java)
	variablesWithInitializers()

	insideFnShortAssign()
	showGolangTypes()
}