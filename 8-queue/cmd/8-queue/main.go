package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/phanty133/id1021/8-queue/pkg/queue"
)

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

func Bench(getQueue func() queue.Queue[int], els []int) (int64, int64) {
	start := time.Now()
	q := getQueue()

	for _, val := range els {
		q.Enqueue(val)
	}

	time1 := time.Since(start).Nanoseconds()
	start = time.Now()

	for range els {
		q.Dequeue()
	}

	time2 := time.Since(start).Nanoseconds()

	return time1, time2
}

func main() {
	sizes := []int{100, 1000, 5000, 10000, 50000, 100000}
	repeats := 100
	// dynamicInitSize := 4
	// timesDynamic1 := make([][]int64, len(sizes))
	// timesStatic1 := make([][]int64, len(sizes))
	timesLL1 := make([][]int64, len(sizes))
	// timesDynamic2 := make([][]int64, len(sizes))
	// timesStatic2 := make([][]int64, len(sizes))
	timesLL2 := make([][]int64, len(sizes))

	for i, size := range sizes {
		fmt.Printf("Running benchmark for size %d\n", size)

		// timesDynamic1[i] = make([]int64, repeats)
		// timesStatic1[i] = make([]int64, repeats)
		timesLL1[i] = make([]int64, repeats)
		// timesDynamic2[i] = make([]int64, repeats)
		// timesStatic2[i] = make([]int64, repeats)
		timesLL2[i] = make([]int64, repeats)

		for j := 0; j < repeats; j++ {
			els := GenRandIntArr(size, 1, 1000000)

			// timesStatic1[i][j], timesStatic2[i][j] = Bench(func() queue.Queue[int] {
			// 	return queue.NewQueueStatic[int](size)
			// }, els)

			// timesDynamic1[i][j], timesDynamic2[i][j] = Bench(func() queue.Queue[int] {
			// 	return queue.NewQueueDynamic[int](dynamicInitSize)
			// }, els)

			timesLL1[i][j], timesLL2[i][j] = Bench(func() queue.Queue[int] {
				return queue.NewQueueLL[int]()
			}, els)

			if j%10 == 0 {
				fmt.Printf("Completed %d/%d\n", j, repeats)
			}
		}
	}

	// WriteTimes("static1.csv", timesStatic1, sizes)
	// WriteTimes("dynamic1.csv", timesDynamic1, sizes)
	WriteTimes("ll1-slow.csv", timesLL1, sizes)
	// WriteTimes("static2.csv", timesStatic2, sizes)
	// WriteTimes("dynamic2.csv", timesDynamic2, sizes)
	WriteTimes("ll2-slow.csv", timesLL2, sizes)
}
