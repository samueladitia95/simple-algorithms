package fileSystem

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadStr() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func ReadInt() int {
	result, err := strconv.Atoi(ReadStr())

	if err != nil {
		log.Fatal(err)
	}

	return result
}
