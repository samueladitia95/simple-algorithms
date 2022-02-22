package gradeStudents

import (
	"errors"
)

func GradeStudents(currentScore int) (string, error) {
	scores := []int{80, 65, 50, 35}
	grades := []string{"A", "B", "C", "D", "E"}

	if currentScore < 0 || currentScore > 100 {
		return "", errors.New("score is invalid")
	}

	for i := 0; i < len(scores); i++ {
		if currentScore >= scores[i] {
			return grades[i], nil
		}
	}
	return grades[len(grades)-1], nil
}
