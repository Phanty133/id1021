package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"

	"github.com/phanty133/id1021/9-heap/pkg/pqueue"
	// "time"
)

/*
Benchmark LL implementations:
- N adds
- N removes (with a pre-allocated queue)

Benchmark Heap:
- N adds
- N removes (with a pre-allocated queue)
- add 1023 rand values, push random increment, measure depth
- Same as 3, but dequeue/queue instead of push

Benchmark Array heap implementations:
- N adds
- N removes (with a pre-allocated queue)
*/

func GenRandIntArr(n, min, max int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max-min) + min
	}

	return arr
}

func WriteTimes(name string, times [][]int64, sizes []int) {
	arrFile, err := os.Create(name)

	if err != nil {
		panic(err)
	}

	defer arrFile.Close()

	arrWriter := csv.NewWriter(arrFile)
	header := make([]string, len(sizes)+1)
	header[0] = "Size"

	for i, size := range sizes {
		header[i+1] = fmt.Sprintf("%d", size)
	}

	if err := arrWriter.Write(header); err != nil {
		panic(err)
	}

	for i := 0; i < len(times[0]); i++ {
		row := make([]string, len(sizes)+1)
		row[0] = fmt.Sprintf("%d", i)

		for j := range sizes {
			row[j+1] = fmt.Sprintf("%d", times[j][i])
		}

		if err := arrWriter.Write(row); err != nil {
			panic(err)
		}
	}

	arrWriter.Flush()
}

/*
Benchmark 1:
- Allocate
- Enqueue N elements
Benchmark 2 (Pre-allocate and enqueue N elements):
- Dequeue N elements
*/

// func Bench(getQueue func() queue.Queue[int], els []int) (int64, int64) {
// 	start := time.Now()
// 	q := getQueue()

// 	for _, val := range els {
// 		q.Enqueue(val)
// 	}

// 	time1 := time.Since(start).Nanoseconds()
// 	start = time.Now()

// 	for range els {
// 		q.Dequeue()
// 	}

// 	time2 := time.Since(start).Nanoseconds()

// 	return time1, time2
// }

func main() {
	h := pqueue.NewArrHeap[int]()
	h.Add(69, 5)
	h.Add(420, 1)
	h.Add(1337, -1)
	h.Add(1, 3)

	for v := range pqueue.IterPQueue[int](h) {
		fmt.Println(v)
	}

	// fmt.Printf("Depth: %d\n", depth)
}
