package main

import (
	"cmp"
	"encoding/csv"
	"fmt"
	"github.com/phanty133/id1021/2-sorted/pkg/bench"
	"os"
)

func NaiveSearch[T cmp.Ordered](arr []T, target T) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}

func NaiveSortedSearch[T cmp.Ordered](arr []T, target T) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] > target {
			return -1
		}

		if arr[i] == target {
			return i
		}
	}

	return -1
}

func BinarySearch[T cmp.Ordered](arr []T, target T) int {
	first := 0
	last := len(arr) - 1

	for {
		idx := (first + last) / 2
		val := arr[idx]

		if val == target {
			return idx
		}

		if val < target && idx < last {
			first = idx + 1
		} else if val > target && idx > first {
			last = idx - 1
		} else {
			return -1
		}
	}
}

func DuplicatesNaive[T cmp.Ordered](a []T, b []T) []T {
	res := make([]T, 0, len(a)+len(b))

	for _, v := range a {
		if NaiveSearch(b, v) != -1 {
			res = append(res, v)
		}
	}

	return res
}

func SortedDuplicatesBinary[T cmp.Ordered](a []T, b []T) []T {
	res := make([]T, 0, len(a)+len(b))

	for _, v := range a {
		if BinarySearch(b, v) != -1 {
			res = append(res, v)
		}
	}

	return res
}

func SortedDuplicatesSmart[T cmp.Ordered](a []T, b []T) []T {
	res := make([]T, 0, len(a)+len(b))

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			res = append(res, a[i])
			i++
			j++
		}
	}

	return res
}

func main() {
	fmt.Println("--- NaiveSearch")
	naiveSearchTimes := bench.BenchSearch(NaiveSearch)

	fmt.Println("--- NaiveSortedSearch")
	naiveSortedTimes := bench.BenchSearch(NaiveSortedSearch)

	fmt.Println("--- BinarySearch")
	binarySearchTimes := bench.BenchSearch(BinarySearch)

	fmt.Println("--- DuplicatesNaive")
	naiveDuplicatesTimes := bench.BenchDuplicates(DuplicatesNaive)

	fmt.Println("--- SortedDuplicatesBinary")
	binaryDuplicatesTimes := bench.BenchDuplicates(SortedDuplicatesBinary)

	fmt.Println("--- SortedDuplicatesSmart")
	smartDuplicatesTimes := bench.BenchDuplicates(SortedDuplicatesSmart)

	outFile, err := os.Create("out-64M.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	exportTimesFunc := func(times []bench.SizeTime) {
		for _, time := range times {
			writer.Write([]string{
				fmt.Sprintf("%d", time.Size),
				fmt.Sprintf("%f", time.Min_us),
				fmt.Sprintf("%f", time.Max_us),
				fmt.Sprintf("%f", time.Median_us),
				fmt.Sprintf("%f", time.Mean_us),
			})
		}
	}

	writer.Write([]string{"Size", "Min_us", "Max_us", "Median_us", "Mean_us"})

	writer.Write([]string{"NaiveSearch"})
	exportTimesFunc(naiveSearchTimes)

	writer.Write([]string{"NaiveSortedSearch"})
	exportTimesFunc(naiveSortedTimes)

	writer.Write([]string{"BinarySearch"})
	exportTimesFunc(binarySearchTimes)

	writer.Write([]string{"DuplicatesNaive"})
	exportTimesFunc(naiveDuplicatesTimes)

	writer.Write([]string{"SortedDuplicatesBinary"})
	exportTimesFunc(binaryDuplicatesTimes)

	writer.Write([]string{"SortedDuplicatesSmart"})
	exportTimesFunc(smartDuplicatesTimes)
}
