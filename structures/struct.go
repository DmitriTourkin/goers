package main

import "fmt"

// Struct — a collection of fields

type Vertex struct {
	X int 
	Y int 
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{} 
	p = &Vertex{1, 2}
)

func vertexs() {
	fmt.Println(v1, v2, v3, p)
}

func ShowStruct() {
	fmt.Println("Show struct")
	fmt.Println(Vertex{1, 2})

	var v Vertex = Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Pointers to structs — can be written without explicit dereference
	secondVertex := Vertex{1, 2}
	p := &secondVertex
	p.X = 1e9
	fmt.Println(secondVertex)
}