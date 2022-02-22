package graduates

import (
	"fmt"

	fileSystem "github.com/samueladitia95/simple-algorithms/03-readTerminal"
)

func App() {
	fmt.Printf("Please enter your name: ")
	name := fileSystem.ReadStr()

	fmt.Printf("Please enter your score: ")
	score := fileSystem.ReadInt()

	fmt.Printf("Please enter how many times you absent from class: ")
	absent := fileSystem.ReadInt()

	status := "Dropped Out"

	if score >= 70 && absent <= 5 {
		status = "Graduates"
	}

	fmt.Printf("%v Status: %v\n", name, status)
}
