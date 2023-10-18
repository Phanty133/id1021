package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/phanty133/id1021/11-trie/pkg/trie"
	"math/rand"
	"os"
	"time"
)

var DATA_PATH string = "/home/phanty/repos/id1021/11-trie/data/kelly.txt"

func GenRandIntArr(n, min, max int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max-min) + min
	}

	return arr
}

func WriteTimes(name string, times [][]int64, headerLabels []string) {
	arrFile, err := os.Create(name)

	if err != nil {
		panic(err)
	}

	defer arrFile.Close()

	arrWriter := csv.NewWriter(arrFile)
	header := make([]string, len(headerLabels)+1)
	header[0] = "Repeat"

	for i, label := range headerLabels {
		header[i+1] = label
	}

	if err := arrWriter.Write(header); err != nil {
		panic(err)
	}

	for i := 0; i < len(times[0]); i++ {
		row := make([]string, len(headerLabels)+1)
		row[0] = fmt.Sprintf("%d", i)

		for j := range headerLabels {
			row[j+1] = fmt.Sprintf("%d", times[j][i])
		}

		if err := arrWriter.Write(row); err != nil {
			panic(err)
		}
	}

	arrWriter.Flush()
}

func Bench(benchFunc func(), repeats int) []int64 {
	times := make([]int64, repeats)

	for i := 0; i < repeats; i++ {
		start := time.Now()
		benchFunc()
		time1 := time.Since(start).Nanoseconds()

		times[i] = time1
	}

	return times
}

func main() {
	words := make([]string, 0)
	file, err := os.Open(DATA_PATH)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Read by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Create trie
	trieStruct := trie.NewTrie()

	for _, word := range words {
		trieStruct.AddWord(word)
	}

	testSeq := trie.WordToSequence("internet")
	fmt.Printf("Sequence: %s\n", testSeq)
	suggested := trieStruct.Lookup(testSeq)

	for _, w := range suggested {
		// Check if original file contained the word
		fmt.Println(w)
	}
	fmt.Printf("Suggested size: %d\n", len(suggested))
}
