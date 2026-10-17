package main

import (
	"fmt"
	"os"

	"github.com/goswamikush/buildlens/internal/parse"
)

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer f.Close()

	parsedMap, err := parse.Parse(f)

	if err != nil {
		fmt.Println("Parse error: ", err)
		return
	}

	fmt.Println("Parsed map: ", parsedMap)
}
