package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/phanty133/id1021/4-linkedlist/pkg/dllist"
	"github.com/phanty133/id1021/4-linkedlist/pkg/llist"
)

/*
Benchmarks:
- Append LL vs array (Vary initial size of list A, append it to fixed size list B, then switch)
- DLL unlink vs LL unlink for K random cells, vary list length
*/

func PrepareLLData(n int) (*llist.LinkedList[int], []*llist.LinkedListItem[int]) {
	l := llist.New[int]()
	items := make([]*llist.LinkedListItem[int], n)

	for i := 0; i < n; i++ {
		curItem := l.Append(i)
		items[i] = curItem
	}

	return l, items
}

func PrepareDLLData(n int) (*dllist.DoublyLinkedList[int], []*dllist.DoublyLinkedListItem[int]) {
	l := dllist.New[int]()
	items := make([]*dllist.DoublyLinkedListItem[int], n)

	for i := 0; i < n; i++ {
		curItem := l.Add(i)
		items[i] = curItem
	}

	return l, items
}

func PrepareArrayData(n int) []int {
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = i
	}

	return a
}

func GenRandomInts(n int, min int, max int) []int {
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(max-min) + min
	}

	return a
}

func main() {
	dynSize := []int{10, 100, 1000, 5000, 10000, 15000}
	fixedSize := 100
	k := 1000
	repeats := 250

	llTimes1 := make([][]int, len(dynSize))
	arrTimes1 := make([][]int, len(dynSize))
	llTimes2 := make([][]int, len(dynSize))
	arrTimes2 := make([][]int, len(dynSize))
	llTimes3 := make([][]int, len(dynSize))
	dllTimes3 := make([][]int, len(dynSize))

	fmt.Println("Appending to linked list vs array")
	fmt.Println("n\tllist\tarray")

	for sizeIdx, n := range dynSize {
		fmt.Println("n =", n)
		fmt.Println("LinkedList")

		llTimes1[sizeIdx] = make([]int, repeats)
		for i := 0; i < repeats; i++ {
			a, _ := PrepareLLData(n)
			b, _ := PrepareLLData(fixedSize)

			start := time.Now()
			nextItem := a.First()

			for nextItem != nil {
				b.Append(nextItem.Head)
				nextItem = nextItem.Next()
			}

			llTimes1[sizeIdx][i] = int(time.Since(start).Microseconds())

			fmt.Printf("%d\n", i)
		}

		fmt.Println("Array")

		arrTimes1[sizeIdx] = make([]int, repeats)
		for i := 0; i < repeats; i++ {
			a := PrepareArrayData(n)
			b := PrepareArrayData(fixedSize)

			start := time.Now()

			b = append(b, a...)

			arrTimes1[sizeIdx][i] = int(time.Since(start).Microseconds())
			fmt.Printf("%d\n", i)
		}
	}

	for sizeIdx, n := range dynSize {
		fmt.Println("n =", n)
		fmt.Println("LinkedList")

		llTimes2[sizeIdx] = make([]int, repeats)
		for i := 0; i < repeats; i++ {
			a, _ := PrepareLLData(fixedSize)
			b, _ := PrepareLLData(n)

			start := time.Now()
			nextItem := a.First()

			for nextItem != nil {
				b.Append(nextItem.Head)
				nextItem = nextItem.Next()
			}

			llTimes2[sizeIdx][i] = int(time.Since(start).Microseconds())
			fmt.Printf("%d\n", i)
		}

		fmt.Println("Array")

		arrTimes2[sizeIdx] = make([]int, repeats)
		for i := 0; i < repeats; i++ {
			a := PrepareArrayData(fixedSize)
			b := PrepareArrayData(n)

			start := time.Now()

			b = append(b, a...)

			arrTimes2[sizeIdx][i] = int(time.Since(start).Microseconds())
			fmt.Printf("%d\n", i)
		}
	}

	repeats2 := 500

	for sizeIdx, n := range dynSize {
		fmt.Println("n =", n)
		fmt.Println("LinkedList")

		kIdxs := GenRandomInts(k, 0, n)

		llTimes3[sizeIdx] = make([]int, repeats2)
		for i := 0; i < repeats2; i++ {
			a, items := PrepareLLData(n)
			kItems := make([]*llist.LinkedListItem[int], k)

			for j := 0; j < k; j++ {
				kItems[j] = items[kIdxs[j]]
			}

			start := time.Now()

			for _, item := range kItems {
				a.Unlink(item)
				a.Insert(item)
			}

			llTimes3[sizeIdx][i] = int(time.Since(start).Microseconds())
			fmt.Printf("%d\n", i)
		}

		fmt.Println("DoublyLinkedList")

		dllTimes3[sizeIdx] = make([]int, repeats2)
		for i := 0; i < repeats2; i++ {
			a, items := PrepareDLLData(n)
			kItems := make([]*dllist.DoublyLinkedListItem[int], k)

			for j := 0; j < k; j++ {
				kItems[j] = items[kIdxs[j]]
			}

			start := time.Now()

			for _, item := range kItems {
				a.Unlink(item)
				a.Insert(item)
			}

			dllTimes3[sizeIdx][i] = int(time.Since(start).Microseconds())
			fmt.Printf("%d\n", i)
		}
	}

	// Write the results to a CSV file
	csvFile, err := os.Create("results.csv")
	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{"", "10", "100", "1000", "5000", "10000", "15000", "10", "100", "1000", "5000", "10000", "15000"})
	writer.Write([]string{"Append to linked list vs array"})
	writer.Write([]string{"n", "llist", "array"})

	for i := 0; i < repeats; i++ {
		row := make([]string, 2*len(dynSize)+1)
		row[0] = fmt.Sprintf("%d", i)

		for j := 0; j < len(dynSize); j++ {
			row[j+1] = fmt.Sprintf("%d", llTimes1[j][i])
			row[j+len(dynSize)+1] = fmt.Sprintf("%d", arrTimes1[j][i])
		}

		writer.Write(row)
	}

	writer.Write([]string{"\n"})
	writer.Write([]string{"Append to linked list vs array"})
	writer.Write([]string{"n", "llist", "array"})

	for i := 0; i < repeats; i++ {
		row := make([]string, 2*len(dynSize)+1)
		row[0] = fmt.Sprintf("%d", i)

		for j := 0; j < len(dynSize); j++ {
			row[j+1] = fmt.Sprintf("%d", llTimes2[j][i])
			row[j+len(dynSize)+1] = fmt.Sprintf("%d", arrTimes2[j][i])
		}

		writer.Write(row)
	}

	writer.Write([]string{"\n"})
	writer.Write([]string{"unlink and insert"})
	writer.Write([]string{"n", "llist", "dllist"})

	for i := 0; i < repeats2; i++ {
		row := make([]string, 2*len(dynSize)+1)
		row[0] = fmt.Sprintf("%d", i)

		for j := 0; j < len(dynSize); j++ {
			row[j+1] = fmt.Sprintf("%d", llTimes3[j][i])
			row[j+len(dynSize)+1] = fmt.Sprintf("%d", dllTimes3[j][i])
		}

		writer.Write(row)
	}

	writer.Write([]string{"\n"})
}
