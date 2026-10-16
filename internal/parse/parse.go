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
				if node.Detail != "" {
					currNode.Detail = node.Detail
				}

				if node.Started != "" {
					currNode.Started = node.Started
				}

				if node.Completed != "" {
					currNode.Completed = node.Completed
				}

				if node.Cached {
					currNode.Cached = node.Cached
				}

				digestMap[node.Digest] = node
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


