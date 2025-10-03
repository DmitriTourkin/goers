package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var (
	m map[string]Vertex
	mapping map[string]Vertex
)


func mapsLiterals() {
	mapping = map[string]Vertex{
		"Yandex" : Vertex{
			40.87865, 56.2453,
		},
		"Google": Vertex{
			67.473475, 83.24387,
		},
	}
	fmt.Println(mapping)
}

func maps() {
	m = make(map[string]Vertex)
	m["HSE University"] = Vertex{
		40.678, -74.8765,
	}
	fmt.Println(m["HSE University"])

	elem1, ok1 := m["Dope"]
	elem2, ok2 := m["HSE University"]
	fmt.Println("Element 1: ", elem1, ok1) // {0 0} false
	fmt.Println("Element 2: ", elem2, ok2) //{40.678 -74.8765} true
}

func main() {
	maps()
	mapsLiterals()
}