package main

import (
	"fmt"
	"math/cmplx"
)

// constants can be character, string, boolean, numberic
const Pi = 3.14

var c, python, java bool

// Basic Types
var (
	ToBe bool = false
	line string = "Hello!"
	rubles float32 = 56.7
	rubles64 float64 = 1024.5
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func showVariables() {
	fmt.Println(rubles64)
	fmt.Println(rubles)
	fmt.Println(MaxInt)
	fmt.Println(z)
}

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
	showVariables()

	typesConversion()
	constantsFn()
}

func typesConversion() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)
}

func constantsFn() {
	const World = "Мир"
	fmt.Println("Hello", World)
	fmt.Println("Happy", "Pi", "Day")
}



