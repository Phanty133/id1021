package main

import (
	"encoding/csv"
	"fmt"
	"github.com/phanty133/id1021/6-trees/pkg/tree"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
Benchmarks:
- Lookup and add exec time for growing node count (dont construct with ordered keys plox)
*/

type TimeEntry struct {
	Lookup int
	Add    int
}

func BenchmarkTree() {
	sizes := []int{100, 1000, 5000, 10000, 50000, 100000, 500000, 1000000}
	adds := 1000
	repeats := 250

	file, err := os.Create("bench.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	cols := make([]string, len(sizes)*2+1)

	cols[0] = "repeat"
	for i, size := range sizes {
		cols[i+1] = "lookup-" + strconv.Itoa(size)
		cols[i+1+len(sizes)] = "add-" + strconv.Itoa(size)
	}

	writer.Write(cols)
	times := make([][]TimeEntry, len(sizes))

	for sizeIdx, size := range sizes {
		times[sizeIdx] = make([]TimeEntry, repeats)

		for i := 0; i < repeats; i++ {
			tree := tree.NewBinaryTree[int, int]()

			for i := 0; i < size; i++ {
				tree.Add(rand.Int(), rand.Int())
			}

			entry := TimeEntry{}
			// addStart := time.Now()

			// for i := 0; i < adds; i++ {
			// 	key := rand.Int()
			// 	tree.Add(key, key)
			// }

			// addTime := float32(time.Since(addStart).Nanoseconds()) / 1000
			lookupStart := time.Now()

			for i := 0; i < adds; i++ {
				tree.Lookup(rand.Int())
			}

			lookupTime := float32(time.Since(lookupStart).Nanoseconds()) / 1000

			entry.Lookup = int(lookupTime)
			entry.Add = int(0)
			times[sizeIdx][i] = entry

			if i%100 == 0 {
				fmt.Printf("Finished repeat %d\n", i)
			}
		}

		fmt.Printf("Finished size %d\n", size)
	}

	for i := 0; i < repeats; i++ {
		row := make([]string, len(sizes)*2+1)
		row[0] = strconv.Itoa(i)

		for sizeIdx, _ := range sizes {
			row[sizeIdx+1] = strconv.Itoa(times[sizeIdx][i].Lookup)
			row[sizeIdx+1+len(sizes)] = strconv.Itoa(times[sizeIdx][i].Add)
		}

		writer.Write(row)
	}
}

func main() {
	BenchmarkTree()

	// tree := tree.NewBinaryTree[int, int]()

	// tree.Add(5, 5)
	// tree.Add(3, 3)
	// tree.Add(7, 7)
	// tree.Add(2, 2)

	// for node := range tree.DepthFirstIterator() {
	// 	fmt.Printf("%d ", node.Val)
	// }
}
