package gradeStudents

import (
	"fmt"
	"log"
	"strconv"

	fileSystem "github.com/samueladitia95/simple-algorithms/03-readTerminal"
)

func App() {
	fmt.Printf(`Please input your current score: `)
	currentScore, err := strconv.Atoi(fileSystem.ReadStr())

	if err != nil {
		log.Fatal(err)
	}

	finalGrade, err := GradeStudents(currentScore)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The Final Score is %v\n", finalGrade)
}
