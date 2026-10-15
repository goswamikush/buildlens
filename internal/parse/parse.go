package parse

import (
	"bufio"
	"fmt"
	"io"
)

func Parse(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("LNIE:", line)
	}

	return scanner.Err()
}


