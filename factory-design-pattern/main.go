package main

import (
	entity "./entity"
)

func main() {
	var student = entity.NewStudent("Rama", "177006099")
	var classroom = entity.NewClassroom(1, "Bunaken")

	student.Greet()
	classroom.ShowClassroom()
}
