package main

import (
	"fmt"
	"strings"
)

func SlicesFunction() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "YYY"
	fmt.Println(a, b) 

	// any changes through slice will be seen, because underlying array is the same

	// [John Paul George Ringo]
// [John Paul] [Paul George]
// [John XXX] [XXX George]
// [John XXX George Ringo]
}

func SliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13} //slice which references this array
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	} {
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
	}

	fmt.Println(s)
}

func SliceLengthAndCapacity() {
	// length — number or elements it contains
	// capacity — the num of element in the underlying array, counting from FRIST el in the slice
	someArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := someArray[3:5]

	fmt.Println(slice)
	fmt.Println("slice len(slice): ", len(slice)) // 2 from index of slice till end of array
	fmt.Println("slice cap(slice): ", cap(slice)) // 7
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
func SlicesWithMake() {
	a := make([]int, 0, 5)
	b := make([]int, 10, 20)
	printSlice("a", a)
	printSlice("b", b)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s []int
	printSlice("s ", s)

	s = append(s, 0)
	s = append(s, 1)
	printSlice("s ", s)
}

