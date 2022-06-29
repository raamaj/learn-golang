// This is implementation of simple Factory Design Pattern
package entity

import "fmt"

type Student struct {
	Name string
	NIK  string
}

func NewStudent(name string, nik string) Student {
	return Student{
		Name: name,
		NIK:  nik,
	}
}

func (s *Student) Greet() {
	fmt.Println("Hello, " + s.Name + "!")
}
