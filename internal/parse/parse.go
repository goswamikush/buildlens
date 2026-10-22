package parse

import (
	"bufio"
	"encoding/json"
	"io"
	"regexp"
	"time"
)

type Node struct {
	Detail    string   `json:"name"`
	Digest    string   `json:"digest"`
	Started   string   `json:"started"`
	Completed string   `json:"completed"`
	Cached    bool     `json:"cached"`
	Inputs    []string `json:"inputs"`
	Status    string
	Duration  time.Duration
	ExtraDetails Details
}

type Details struct {
	Step string
	Total string
	Stage string
	Keyword string
	Instruction string
}

type vertex struct {
	Nodes []Node `json:"vertexes"`
}

var buildRegex = regexp.MustCompile(`^\[(?:([^\s\]]+)\s)?(\d+)\/(\d+)\]\s+((\S+).*)$`)

func Parse(r io.Reader) (map[string]Node, []string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)

	digestMap := map[string]Node{}
	order := []string{}

	for scanner.Scan() {
		var v vertex
		err := json.Unmarshal(scanner.Bytes(), &v)

		if err != nil {
			continue
		}

		for _, node := range v.Nodes {
			currNode, ok := digestMap[node.Digest]

			if ok {
				node = mergeNodes(node, currNode)
			} else {
				order = append(order, node.Digest)
			}

			digestMap[node.Digest] = node
		}
	}

	if scanner.Err() != nil {
		return nil, nil, scanner.Err() 
	}

	digestMap, err := finalize(digestMap)

	if err != nil {
		return nil, nil, err
	}

	return digestMap, order, nil
}

func finalize(digestMap map[string]Node) (map[string]Node, error) {
	for digest, node := range digestMap {
		newDetails, _ := parseDetails(node.Detail)
		node.ExtraDetails = newDetails

		node.Status = "incomplete"

		if node.Completed != "" {
			node.Status = "executed"
		}

		if node.Cached {
			node.Status = "cached"
		}

		if node.Started == "" || node.Completed == "" {
			digestMap[digest] = node
			continue
		}

		start, err := time.Parse(time.RFC3339Nano, node.Started)

		if err != nil {
			return nil, err
		}

		end, err := time.Parse(time.RFC3339Nano, node.Completed)

		if err != nil {
			return nil, err
		}

		duration := end.Sub(start)

		node.Duration = duration

		digestMap[digest] = node
	}

	return digestMap, nil
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

func parseDetails(input string) (Details, bool) {
	res :=  buildRegex.FindStringSubmatch(input)

	var newDetails Details

	if len(res) >= 2 && res[1] != "" {
		newDetails.Stage = res[1]
	}

	if len(res) >= 3 && res[2] != "" {
		newDetails.Step = res[2]
	}

	if len(res) >= 4 && res[3] != "" {
		newDetails.Total = res[3]
	}

	if len(res) >= 5 && res[4] != "" {
		newDetails.Instruction = res[4]
	}

	if len(res) >= 6 && res[5] != "" {
		newDetails.Keyword = res[5]
	}

	return newDetails, true
}
