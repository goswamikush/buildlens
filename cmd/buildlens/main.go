package main

import (
	"fmt"
	"os"
	"slices"

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

	var nodes []parse.Node
	for _, digest := range order {
		newNode := parsedMap[digest]

		if newNode.Internal {
			continue
		}

		nodes = append(nodes, newNode)
	}

	slices.SortFunc(nodes, func(a, b parse.Node) int {
		return a.Started.Compare(b.Started)
	})

	for _, node := range nodes {
		d := node.ExtraDetails

		fmt.Println(d.Stage, d.Step + "/" + d.Total, d.Instruction, node.Status, node.Duration)
	}
}
