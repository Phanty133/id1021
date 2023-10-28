package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/phanty133/id1021/11-trie/pkg/trie"
	"os"
)

var DATA_PATH string = "/home/phanty/repos/id1021/11-trie/data/kelly.txt"

func main() {
	words := make([]string, 0)
	file, err := os.Open(DATA_PATH)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Read by line
	scanner := bufio.NewScanner(file)
	unique := make(map[string]bool, 0)

	for scanner.Scan() {
		w := scanner.Text()

		if unique[w] {
			continue
		}

		words = append(words, w)
		unique[w] = true
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Create trie
	trieStruct := trie.NewTrie()

	for _, word := range words {
		trieStruct.AddWord(word)
	}

	// Count which words have duplicates
	duplicates := make(map[string]int, 0)

	for _, w := range words {
		seq := trie.WordToSequence(w)
		suggested := trieStruct.Lookup(seq)

		for _, s := range suggested {
			duplicates[s]++
		}
	}

	// Write duplicates to csv

	file, err = os.Create("duplicates.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	for word, count := range duplicates {
		writer.Write([]string{word, fmt.Sprintf("%d", count)})
	}

	writer.Flush()
}
