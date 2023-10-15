package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"

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
			if i >= len(times[j]) {
				continue
			}

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
- Add N elements
Benchmark 2 (Pre-allocate and add N elements):
- Remove N elements
*/

func Bench(q pqueue.PriorityQueue[int], els []int) (int64, int64) {
	start := time.Now()

	for _, val := range els {
		q.Add(val, val)
	}

	time1 := time.Since(start).Nanoseconds()
	start = time.Now()

	for range els {
		q.Remove()
	}

	time2 := time.Since(start).Nanoseconds()

	return time1, time2
}

func main() {
	sizes := []int{100, 500, 1000, 2500, 5000, 7500, 10000, 25000, 50000, 75000, 100000}

	repeats := 100
	timesLLFA1 := make([][]int64, len(sizes))
	timesLLFR1 := make([][]int64, len(sizes))
	timesHeap1 := make([][]int64, len(sizes))
	timesArrHeap1 := make([][]int64, len(sizes))
	timesLLFR2 := make([][]int64, len(sizes))
	timesLLFA2 := make([][]int64, len(sizes))
	timesHeap2 := make([][]int64, len(sizes))
	timesArrHeap2 := make([][]int64, len(sizes))

	for i, size := range sizes {
		fmt.Printf("Running benchmark for size %d\n", size)

		timesLLFA1[i] = make([]int64, repeats)
		timesLLFR1[i] = make([]int64, repeats)
		timesHeap1[i] = make([]int64, repeats)
		timesArrHeap1[i] = make([]int64, repeats)
		timesLLFR2[i] = make([]int64, repeats)
		timesLLFA2[i] = make([]int64, repeats)
		timesHeap2[i] = make([]int64, repeats)
		timesArrHeap2[i] = make([]int64, repeats)

		for j := 0; j < repeats; j++ {
			els := GenRandIntArr(size, 1, 1000000)

			timesLLFR1[i][j], timesLLFR2[i][j] = Bench(pqueue.NewPQueueLLFastRemove[int](), els)
			timesLLFA1[i][j], timesLLFA2[i][j] = Bench(pqueue.NewPQueueLLFastAdd[int](), els)
			timesHeap1[i][j], timesHeap2[i][j] = Bench(pqueue.NewHeap[int](), els)
			timesArrHeap1[i][j], timesArrHeap2[i][j] = Bench(pqueue.NewArrHeap[int](0), els)

			if j%10 == 0 {
				fmt.Printf("--- Completed %d/%d\n", j, repeats)

				fmt.Println("Sample times: ")
				fmt.Printf("LLFR: %dus, %dus\n", timesLLFR1[i][j]/1000, timesLLFR2[i][j]/1000)
				fmt.Printf("LLFA: %dus, %dus\n", timesLLFA1[i][j]/1000, timesLLFA2[i][j]/1000)
				fmt.Printf("Heap: %dus, %dus\n", timesHeap1[i][j]/1000, timesHeap2[i][j]/1000)
				fmt.Printf("ArrHeap: %dus, %dus\n", timesArrHeap1[i][j]/1000, timesArrHeap2[i][j]/1000)
			}
		}
	}

	WriteTimes("add-llfr.csv", timesLLFR1, sizes)
	WriteTimes("add-llfa.csv", timesLLFA1, sizes)
	WriteTimes("add-heap-1.csv", timesHeap1, sizes)
	WriteTimes("add-arrheap.csv", timesArrHeap1, sizes)
	WriteTimes("remove-llfr.csv", timesLLFR2, sizes)
	WriteTimes("remove-llfa.csv", timesLLFA2, sizes)
	WriteTimes("remove-heap-1.csv", timesHeap2, sizes)
	WriteTimes("remove-arrheap.csv", timesArrHeap2, sizes)

	// Push/Add-Remove benchmark

	repeats = 500
	// pushHeapSize := 1023
	timesPush := make([][]int64, len(sizes))
	timesAdd := make([][]int64, len(sizes))
	// pushDepth := make([][]int64, len(sizes))

	for i, size := range sizes {
		fmt.Printf("Running benchmark for size %d\n", size)

		timesPush[i] = make([]int64, repeats)
		timesAdd[i] = make([]int64, repeats)

		for j := 0; j < repeats; j++ {
			els := GenRandIntArr(size, 1, 10000)
			q1 := pqueue.NewHeap[int]()
			q2 := pqueue.NewHeap[int]()

			for _, val := range els {
				q1.Add(val, val)
				q2.Add(val, val)
			}

			start := time.Now()

			for k := 0; k < repeats; k++ {
				q1.Push(rand.Intn(10000))
				// pushDepth[i] = append(pushDepth[i], int64(d))
			}

			time1 := time.Since(start).Nanoseconds()
			start = time.Now()

			for k := 0; k < size; k++ {
				val, _ := q2.Remove()
				q2.Add(val, rand.Intn(100))
			}

			time2 := time.Since(start).Nanoseconds()

			timesPush[i][j] = time1
			timesAdd[i][j] = time2
		}
	}

	WriteTimes("push.csv", timesPush, sizes)
	WriteTimes("add-remove.csv", timesAdd, sizes)
	// WriteTimes("push-depth.csv", pushDepth, sizes)
}
