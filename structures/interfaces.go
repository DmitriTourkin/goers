package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// About structs
type Employee struct {
	name string
	age int
	isRemote bool
}

type Address struct {
	Street string `json:"currentCity"`
	City string `json:"myCity"`
}

type YetAnotherEmployee struct {
	Name string `json:"name"`
	Age int 		`json:"age"`
	IsBusy bool `json:"isBusy"`
	Address // composition
}

func (e *Employee) updateName(newName string) {
	e.name = newName
}

func main() {
	employee1 := Employee{
		name: "Dima",
		age: 21, 
		isRemote: true,
	}

	employee1.updateName("Dmitry")

	// anonymous stucts
	jobs := struct {
		level int
		msg string
	}{
		level: 56,
		msg: "Differents",
	}

	fmt.Println(jobs)

	employee2 := YetAnotherEmployee{"Kate", 56, false, Address{"Yellow Street", "Moscow"}}

	jsonData, _ := json.Marshal(employee2)
	fmt.Println(string(jsonData))
}



// Interfaces
type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius
}

func calculateShape(s Shape) float64 {
	return s.Area()
}

func describeValue(t interface{}) { // VERY BROAD IMPLEMENTATION
	fmt.Printf("Type %T, Value: %v", t, t)
}






