package fileSystem

import (
	"bufio"
	"os"
)

func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}
