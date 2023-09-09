// Reimplementation and improvement upon the Bench.java file provided in the course

package bench

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func CreatedSortedIntArray(n int, max_delta int) []int {
	arr := make([]int, n)
	curVal := 0

	for i := 0; i < n; i++ {
		curVal += rand.Intn(max_delta) + 1
		arr[i] = curVal
	}

	return arr
}

func keys(loop, n int) []int {
	indx := make([]int, loop)

	for i := 0; i < loop; i++ {
		indx[i] = rand.Intn(n * 5)
	}

	return indx
}

type SizeTime struct {
	Size      int
	Min_us    float64
	Max_us    float64
	Median_us float64
	Mean_us   float64
}

func median(arr []float64) float64 {
	n := len(arr)
	mid := n / 2

	if n%2 == 0 {
		return (arr[mid-1] + arr[mid]) / 2
	} else {
		return arr[mid]
	}
}

func mean(arr []float64) float64 {
	sum := 0.0

	for _, val := range arr {
		sum += val
	}

	return sum / float64(len(arr))
}

func BenchSearch(search func([]int, int) int) []SizeTime {
	sizes := []int{
		100, 200, 300, 400, 500,
		600, 700, 800, 900, 1000,
		1100, 1200, 1300, 1400,
		1500, 1600, 10000,
		100000, 1000000,
	}
	times := make([]SizeTime, 0, len(sizes))

	for _, size := range sizes {
		loop := 10000
		sortedArr := CreatedSortedIntArray(size, 1000)
		indx := keys(loop, size)

		fmt.Printf("Size: %d\n", size)

		iters := 1000
		minTime := math.Inf(1)
		maxTime := math.Inf(-1)
		iterTimes := make([]float64, iters)

		for i := 0; i < iters; i++ {
			start := time.Now()

			for _, i := range indx {
				search(sortedArr, i)
			}

			elapsed := time.Since(start).Nanoseconds()
			minTime = math.Min(minTime, float64(elapsed))
			maxTime = math.Max(maxTime, float64(elapsed))
			iterTimes[i] = float64(elapsed)

			if i%100 == 0 {
				fmt.Printf("Iteration %d. Last elapsed time: %fus\n", i, float64(elapsed)/1000)
			}
		}

		times = append(times, SizeTime{
			size,
			minTime / 1000,
			maxTime / 1000,
			median(iterTimes) / 1000,
			mean(iterTimes) / 1000,
		})
	}

	return times
}

func BenchDuplicates(findDuplicates func([]int, []int) []int) []SizeTime {
	sizes := []int{
		100, 200, 300, 400, 500,
		600, 700, 800, 900, 1000,
		1100, 1200, 1300, 1400,
		1500, 1600, 10000,
		100000, 1000000,
	}
	times := make([]SizeTime, 0, len(sizes))

	for _, size := range sizes {
		sortedArrA := CreatedSortedIntArray(size, 1000)
		sortedArrB := CreatedSortedIntArray(size, 1000)

		fmt.Printf("Size: %d\n", size)

		iters := 1000
		minTime := math.Inf(1)
		maxTime := math.Inf(-1)
		iterTimes := make([]float64, iters)

		for i := 0; i < iters; i++ {
			start := time.Now()

			findDuplicates(sortedArrA, sortedArrB)

			elapsed := time.Since(start).Nanoseconds()
			minTime = math.Min(minTime, float64(elapsed))
			maxTime = math.Max(maxTime, float64(elapsed))
			iterTimes[i] = float64(elapsed)

			if i%100 == 0 {
				fmt.Printf("Iteration %d. Last elapsed time: %fus\n", i, float64(elapsed)/1000)
			}
		}

		times = append(times, SizeTime{
			size,
			minTime / 1000,
			maxTime / 1000,
			median(iterTimes) / 1000,
			mean(iterTimes) / 1000,
		})
	}

	return times
}
