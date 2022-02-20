package main

import (
	"fmt"
	"log"
)

func main() {
	currentScore := 10

	err, finalGrade := GradeStudents(currentScore)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The Final Score is %v\n", finalGrade)
}
