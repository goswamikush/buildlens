package parse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Node struct {
	Detail string `json:"name"`
	Digest string `json:"digest"`
	Started string `json:"started"` 
	Completed string `json:"completed"`
	Cached bool `json:"cached"`
	Inputs []string `json:"inputs"`
	Status string
	Duration time.Duration
}

type vertex struct {
	Nodes []Node `json:"vertexes"`
}

func Parse(r io.Reader) (map[string]Node, error) {
	scanner := bufio.NewScanner(r)

	digestMap := map[string]Node{}

	for scanner.Scan() {
		var v vertex;
		err := json.Unmarshal(scanner.Bytes(), &v);

		if err != nil {
			continue
		}

		for _, node := range v.Nodes {
			currNode, ok := digestMap[node.Digest] 

			if ok {
				node = mergeNodes(node, currNode)
			}

			digestMap[node.Digest] = node
		}
	}

	digestMap = finalize(digestMap)

	return digestMap, scanner.Err()
}

func finalize(digestMap map[string]Node) (map[string]Node){
	for digest, node := range digestMap {
		start, err := time.Parse(time.RFC3339Nano, node.Started)

		if err != nil {
			fmt.Println("Failed to parse timestamp: ", err)
		}

		end, err := time.Parse(time.RFC3339Nano, node.Completed)

		if err != nil {
			fmt.Println("Failed to parse timestamp: ", err)
		}

		duration := end.Sub(start)

		fmt.Println(duration)

		node.Duration = duration

		digestMap[digest] = node
	}

	print(digestMap)

	return digestMap
}

func mergeNodes(newNode Node, oldNode Node) Node {
	if newNode.Detail != "" {
		oldNode.Detail = newNode.Detail
	}

	if newNode.Started != "" {
		oldNode.Started = newNode.Started
	}

	if newNode.Completed != "" {
		oldNode.Completed = newNode.Completed
	}

	if len(newNode.Inputs) != 0 {
		oldNode.Inputs = newNode.Inputs
	}

	if newNode.Cached {
		oldNode.Cached = newNode.Cached
	}

	return oldNode
}


