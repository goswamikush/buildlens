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

	if err := parse.Parse(f); err != nil {
		fmt.Println("parse error:", err)
	}
}
