package graduates

import "fmt"

func App() {
	name := "Samuel Aditia"
	score := 67
	absent := 4
	status := "Dropped Out"

	if score >= 70 && absent <= 5 {
		status = "Graduates"
	}

	fmt.Printf("%v Status: %v\n", name, status)
}
