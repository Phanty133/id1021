package postnum

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Code    string
	CodeNum int
	Name    string
	Pop     int
}

func ZipToNum(zip string) (int, error) {
	// Converts the given zip code to an integer. If the zip code is invalid,
	// an error is returned.

	zipNum, err := strconv.Atoi(strings.ReplaceAll(zip, " ", ""))

	if err != nil {
		return 0, fmt.Errorf("error converting zipcode to int (Given %s)", zip)
	}

	return zipNum, nil
}

func ReadData(path string) []Node {
	// Reads in the CSV file at path and returns a slice of Node structs
	// representing the data in the file.

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	nodes := make([]Node, len(records))

	for i, record := range records {
		pop, err := strconv.Atoi(record[2])

		if err != nil {
			fmt.Printf("Error converting population to int (Given %s)\n", record[2])
			continue
		}

		codeNum, err := ZipToNum(record[0])

		if err != nil {
			fmt.Println(err)
			continue
		}

		nodes[i] = Node{
			Code:    strings.TrimSpace(record[0]),
			CodeNum: codeNum,
			Name:    strings.TrimSpace(record[1]),
			Pop:     pop,
		}
	}

	return nodes
}

func FindZipCodeLinear(data []Node, zip string) *Node {
	// Finds the Node with the given zip code in the slice of Nodes using a
	// linear search. If no Node is found, nil is returned.

	for _, node := range data {
		if node.Code == zip {
			return &node
		}
	}

	return nil
}

func FindZipCodeLinearNum(data []Node, zip int) *Node {
	// Finds the Node with the given zip code in the slice of Nodes using a
	// linear search. If no Node is found, nil is returned.

	for _, node := range data {
		if node.CodeNum == zip {
			return &node
		}
	}

	return nil
}

func FindZipCodeBinary(data []Node, zip string) *Node {
	// Finds the Node with the given zip code in the slice of Nodes using a
	// binary search. If no Node is found, nil is returned.

	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2

		if data[mid].Code == zip {
			return &data[mid]
		} else if data[mid].Code < zip {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return nil
}

func FindZipCodeBinaryNum(data []Node, zip int) *Node {
	// Finds the Node with the given zip code in the slice of Nodes using a
	// binary search. If no Node is found, nil is returned.

	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2

		if data[mid].CodeNum == zip {
			return &data[mid]
		} else if data[mid].CodeNum < zip {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return nil
}

func CreatePlainIndexedMap(data []Node) []Node {
	// Creates a plain indexed map from the slice of Nodes. The map is indexed
	// by the zip code.

	maxZip := data[0].CodeNum

	for _, node := range data {
		if node.CodeNum > maxZip {
			maxZip = node.CodeNum
		}
	}

	m := make([]Node, maxZip+1)

	for _, node := range data {
		m[node.CodeNum] = node
	}

	return m
}

func LookupPlainIndex(m []Node, zip string) *Node {
	// Looks up the Node with the given zip code in the plain indexed map. If
	// no Node is found, nil is returned.

	zipNum, err := ZipToNum(zip)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &m[zipNum]
}
