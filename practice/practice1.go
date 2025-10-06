package main

import (
	"fmt"
)

type Person struct {
	Id int
	FirstName string
	LastName string 
	Email string 
}

type Professor struct {
	Person
	Department string
	Salary float64
	CourseIds []int
	IsFullTime bool
}

type Student struct {
	Person
	Major string
	Gpa float64
	CourseIDs []int
	IsGraduated bool
}

type Course struct {
	Id int
	Name string 
	Credits int
	ProfessorId *int
	StudentIds []int
}

type UniversitySystem struct {
	Students map[int]*Student
	Professors map[int]*Professor
	Courses map[int]*Course
	StudentIds []int
	ProfessorIds []int
	CoursesIds []int
}

type UniversityMember interface {
	getInfo() string 
	getRole() string
}

func (s *Student) getInfo() string {
	var status string 

	if s.IsGraduated {
		status = "Выпускник"
	} else {
		status = "Обучающийся"
	}
	return fmt.Sprintf("ФИ: %s %s, GPA: %f, Major: %s, Status: %s", s.FirstName, s.LastName, s.Gpa, s.Major, status)
}

func (p *Professor) getInfo() string {
	var employment string = "Полная ставка"

	if !p.IsFullTime {
		employment = "Частичная ставка"
	} 

	return fmt.Sprintf("ФИ: %s %s, Department: %s, Salary: %f, Employment: %s", p.FirstName, p.LastName, p.Department, p.Salary, employment)
}

func (s *Student) getRole() string {
	return "Student"
}

func (p *Professor) getRole() string {
	return "Professor"
}

func main() {
	system := &UniversitySystem{
		Students: make(map[int]*Student),
		Professors: make(map[int]*Professor), 
		Courses: make(map[int]*Course),
	}

	fmt.Println(*system)
	fmt.Println(&system)
}