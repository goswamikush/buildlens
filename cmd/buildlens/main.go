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

	parsedMap, order, err := parse.Parse(f)

	if err != nil {
		fmt.Println("Parse error: ", err)
		return
	}

	// for _, node := range parsedMap {
	// 	fmt.Println(node.Detail, node.Status, node.Duration)
	// }

	for _, digest := range order {
		fmt.Println(parsedMap[digest])
	}
}
