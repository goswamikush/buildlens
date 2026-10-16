package parse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type node struct {
	Detail string `json:"name"`
	Digest string `json:"digest"`
	Started string `json:"started"` 
	Completed string `json:"completed"`
	Cached bool `json:"cached"`
}

type vertex struct {
	Nodes []node `json:"vertexes"`
}

func Parse(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	digestMap := map[string]node{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("LINE:", line)

		var v vertex;
		err := json.Unmarshal(scanner.Bytes(), &v);

		for _, node := range v.Nodes {
			currNode, ok := digestMap[node.Digest] 

			if ok {
				node = mergeNodes(node, currNode)
			} 
			digestMap[node.Digest] = node
		}

		if err != nil {
			continue
		}

		fmt.Println("Parsed:")
		fmt.Println(v)
		fmt.Println()

	}

	fmt.Println(digestMap)

	return scanner.Err()
}

func mergeNodes(newNode node, oldNode node) node {
	if newNode.Detail != "" {
		oldNode.Detail = newNode.Detail
	}

	if newNode.Started != "" {
		oldNode.Started = newNode.Started
	}

	if newNode.Completed != "" {
		oldNode.Completed = newNode.Completed
	}

	if newNode.Cached {
		oldNode.Cached = newNode.Cached
	}

	return oldNode
}


