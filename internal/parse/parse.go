package parse

import (
	"bufio"
	"encoding/json"
	"io"
)

type Node struct {
	Detail string `json:"name"`
	Digest string `json:"digest"`
	Started string `json:"started"` 
	Completed string `json:"completed"`
	Cached bool `json:"cached"`
	Inputs []string `json:"inputs"`
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

	return digestMap, scanner.Err()
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


