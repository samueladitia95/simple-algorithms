package gradeStudents

import (
	"fmt"
	"log"
)

func App() {
	currentScore := 10

	finalGrade, err := GradeStudents(currentScore)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The Final Score is %v\n", finalGrade)
}
