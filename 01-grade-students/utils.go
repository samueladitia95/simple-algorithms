package main

import (
	"errors"
)

func GradeStudents(currentScore int) (error, string) {
	scores := []int{80, 65, 50, 35}
	grades := []string{"A", "B", "C", "D", "E"}

	if currentScore < 0 || currentScore > 100 {
		return errors.New("Score is Invalid"), ""
	}

	for i := 0; i < len(scores); i++ {
		if currentScore >= scores[i] {
			return nil, grades[i]
		}
	}
	return nil, grades[len(grades)-1]
}
