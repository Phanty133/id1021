package postnum

import (
	"fmt"
)

type NodeHashMap struct {
	buckets    [][]Node
	numBuckets int
}

func CreateHashMap(data []Node, numBuckets int) *NodeHashMap {
	// Creates a hash map with the given number of buckets and inserts the
	// given slice of Nodes into it.

	hashMap := NodeHashMap{
		buckets:    make([][]Node, numBuckets),
		numBuckets: numBuckets,
	}

	for _, node := range data {
		hash := node.CodeNum % numBuckets

		if hashMap.buckets[hash] == nil {
			hashMap.buckets[hash] = make([]Node, 0)
		}

		hashMap.buckets[hash] = append(hashMap.buckets[hash], node)
	}

	return &hashMap
}

func (data *NodeHashMap) Lookup(zip string) (*Node, int64) {
	// Looks up the Node with the given zip code in the hash map. If no Node is
	// found, nil is returned.
	// Returns size of bucket the Node was found in.

	zipNum, err := ZipToNum(zip)

	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	hash := zipNum % data.numBuckets

	for _, node := range data.buckets[hash] {
		if node.Code == zip {
			return &node, int64(len(data.buckets[hash]))
		}
	}

	return nil, 0
}

func CountCollisions(data []Node, m int) int64 {
	// Counts the number of collisions that would occur if the slice of Nodes
	// was inserted into a hash map with the given size.

	var count int64 = 0
	hashMap := make([]Node, m)

	for _, node := range data {
		hash := node.CodeNum % m

		if hashMap[hash].CodeNum != 0 {
			count++
		}

		hashMap[hash] = node
	}

	return count
}
