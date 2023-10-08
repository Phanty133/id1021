package main

import (
	"encoding/csv"
	"fmt"
	"github.com/phanty133/id1021/7-quicksort/pkg/llist"
	"math/rand"
	"os"
	"time"
)

func GenRandIntArr(n, min, max int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max-min) + min
	}

	return arr
}

func ArrToLinkedList(arr []int) *llist.LinkedList[int] {
	list := llist.New[int]()

	for _, v := range arr {
		list.Append(v)
	}

	return list
}

func LinkedListToArr(ll *llist.LinkedList[int]) []int {
	arr := make([]int, ll.Length())
	i := 0

	for item := ll.First(); item != nil; item = item.Next() {
		arr[i] = item.Head
		i++
	}

	return arr
}

func partitionArray(array []int, min, max int) int {
	pivot := array[min]
	i := min
	j := max

	for i < j {
		for array[i] <= pivot && i < max {
			i++
		}

		for array[j] > pivot && j > min {
			j--
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	array[min], array[j] = array[j], array[min]

	return j
}

func QuickSortArrayInPlace(array []int) {
	quickSortArrayInPlaceRec(array, 0, len(array)-1)
}

func quickSortArrayInPlaceRec(array []int, min, max int) {
	if min >= max {
		return
	}

	pivot := partitionArray(array, min, max)

	quickSortArrayInPlaceRec(array, min, pivot-1)
	quickSortArrayInPlaceRec(array, pivot+1, max)
}

func QuickSortLinkedList(list *llist.LinkedList[int]) {
	quickSortLinkedListRec(list.First(), list.Last())
}

func quickSortLinkedListRec(min, max *llist.LinkedListItem[int]) {
	if min == max {
		return
	}

	pivot := partitionLinkedList(min, max)

	if pivot != nil {
		if pivot.Next() != nil {
			quickSortLinkedListRec(pivot.Next(), max)
		}

		if pivot != min {
			quickSortLinkedListRec(min, pivot)
		}
	}
}

func partitionLinkedList(min, max *llist.LinkedListItem[int]) *llist.LinkedListItem[int] {
	pivot := min
	i := min

	for i != max && i != nil {
		if i.Head < max.Head {
			pivot = min

			i.Head, min.Head = min.Head, i.Head

			min = min.Next()
		}

		i = i.Next()
	}

	min.Head, max.Head = max.Head, min.Head
	return pivot
}

func BenchmarkQuickSortArray(arr []int) int64 {
	start := time.Now()
	QuickSortArrayInPlace(arr)
	return time.Since(start).Nanoseconds()
}

func BenchmarkQuickSortLinkedList(ll *llist.LinkedList[int]) int64 {
	start := time.Now()
	QuickSortLinkedList(ll)
	return time.Since(start).Nanoseconds()
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

		for j, _ := range sizes {
			row[j+1] = fmt.Sprintf("%d", times[j][i])
		}

		if err := arrWriter.Write(row); err != nil {
			panic(err)
		}
	}

	arrWriter.Flush()
}

func main() {
	objSizes := []int{10, 100, 1000, 5000, 10000, 50000, 100000, 500000, 1000000}
	repeats := 500
	arrTimes := make([][]int64, len(objSizes))
	llTimes := make([][]int64, len(objSizes))

	for i, size := range objSizes {
		fmt.Printf("Size: %d\n", size)

		arrTimes[i] = make([]int64, repeats)
		llTimes[i] = make([]int64, repeats)

		for j := 0; j < repeats; j++ {
			arr := GenRandIntArr(size, 0, 1000000)
			ll := ArrToLinkedList(arr)

			arrTimes[i][j] = BenchmarkQuickSortArray(arr)
			llTimes[i][j] = BenchmarkQuickSortLinkedList(ll)

			if j%100 == 0 {
				fmt.Printf("Repeat: %d\n", j)
			}
		}
	}

	WriteTimes("array.csv", arrTimes, objSizes)
	WriteTimes("linkedlist.csv", llTimes, objSizes)
}
