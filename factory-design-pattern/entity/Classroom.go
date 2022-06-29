// This is implementation of interface Factory Design Pattern
package entity

import (
	"fmt"
)

// Create interface for classroom functions
type Classroom interface {
	// Method for showing the classroom information
	ShowClassroom()
}

// Create local classroom struct
type classroom struct {
	RoomNumber int
	ClassName  string
}

// Implementation of ShowClassroom method
func (c classroom) ShowClassroom() {
	fmt.Printf("Your Classroom is %d : %s\n", c.RoomNumber, c.ClassName)
}

// Method for creating new classroom instance
func NewClassroom(roomNumber int, classname string) Classroom {
	return classroom{
		RoomNumber: roomNumber,
		ClassName:  classname,
	}
}
