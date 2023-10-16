package main

import (
	"encoding/csv"
	"fmt"
	"github.com/phanty133/id1021/10-hashmap/pkg/postnum"
	"math/rand"
	"os"
	"time"
)

var DATA_PATH string = "/home/phanty/repos/id1021/10-hashmap/data/postnummer.csv"

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

func Bench(searchFunc func(), repeats int) []int64 {
	times := make([]int64, repeats)

	for i := 0; i < repeats; i++ {
		start := time.Now()
		searchFunc()
		time1 := time.Since(start).Nanoseconds()

		times[i] = time1
	}

	return times
}

func BenchCollisions(minModulo, maxModulo int) []int64 {
	collisions := make([]int64, maxModulo-minModulo)
	data := postnum.ReadData(DATA_PATH)

	for i := minModulo; i < maxModulo; i++ {
		collisions[i-minModulo] = postnum.CountCollisions(data, i)

		if i%100 == 0 {
			fmt.Printf("Finished modulo %d\n", i)
		}
	}

	return collisions
}

// Benchmarks:
// - String linear and binary search for 111 15 and 984 99
// - Int linear and binary search for 111 15 and 984 99
// - vanilla map lookup for 111 15 and 984 99
// - number of collisions histogram
// - lookup time for hashmap of varying sizes
// - bucket sizes for hashmap of varying sizes (histogram)

func main() {
	data := postnum.ReadData(DATA_PATH)
	repeats := 5000

	stringLinearTimes1 := Bench(func() {
		postnum.FindZipCodeLinear(data, "111 15")
	}, repeats)

	stringLinearTimes2 := Bench(func() {
		postnum.FindZipCodeLinear(data, "984 99")
	}, repeats)

	stringBinaryTimes1 := Bench(func() {
		postnum.FindZipCodeBinary(data, "111 15")
	}, repeats)

	stringBinaryTimes2 := Bench(func() {
		postnum.FindZipCodeBinary(data, "984 99")
	}, repeats)

	intLinearTimes1 := Bench(func() {
		postnum.FindZipCodeLinearNum(data, 11115)
	}, repeats)

	intLinearTimes2 := Bench(func() {
		postnum.FindZipCodeLinearNum(data, 98499)
	}, repeats)

	intBinaryTimes1 := Bench(func() {
		postnum.FindZipCodeBinaryNum(data, 11115)
	}, repeats)

	intBinaryTimes2 := Bench(func() {
		postnum.FindZipCodeBinaryNum(data, 98499)
	}, repeats)

	plainMap := postnum.CreatePlainIndexedMap(data)

	plainMapTimes1 := Bench(func() {
		postnum.LookupPlainIndex(plainMap, "111 15")
	}, repeats)

	plainMapTimes2 := Bench(func() {
		postnum.LookupPlainIndex(plainMap, "984 99")
	}, repeats)

	hashMap := postnum.CreateHashMap(data, 10000)

	hashMapTimes1 := Bench(func() {
		hashMap.Lookup("111 15")
	}, repeats)

	hashMapTimes2 := Bench(func() {
		hashMap.Lookup("984 99")
	}, repeats)

	basicBenchTimes := [][]int64{
		stringLinearTimes1,
		stringLinearTimes2,
		stringBinaryTimes1,
		stringBinaryTimes2,
		intLinearTimes1,
		intLinearTimes2,
		intBinaryTimes1,
		intBinaryTimes2,
		plainMapTimes1,
		plainMapTimes2,
		hashMapTimes1,
		hashMapTimes2,
	}

	basicBenchHeader := []string{
		"String linear 111 15",
		"String linear 984 99",
		"String binary 111 15",
		"String binary 984 99",
		"Int linear 111 15",
		"Int linear 984 99",
		"Int binary 111 15",
		"Int binary 984 99",
		"Plain map 111 15",
		"Plain map 984 99",
		"Hash map 111 15",
		"Hash map 984 99",
	}

	WriteTimes("basic_bench.csv", basicBenchTimes, basicBenchHeader)

	collisions := BenchCollisions(1, 100000)

	WriteTimes("collisions.csv", [][]int64{collisions}, []string{"Collisions"})

	hashMapSizes := []int{10, 100, 500, 1000, 2500, 5000, 7500, 10000, 15000, 25000, 37500, 50000, 62500, 75000, 100000}
	hashMapTimesVaried := make([][]int64, len(hashMapSizes))
	hashMapBucketSizes := make([][]int64, len(hashMapSizes))

	lookupIds := GenRandIntArr(repeats, 0, len(data))
	lookupZipCodes := make([]string, repeats)

	for i, id := range lookupIds {
		lookupZipCodes[i] = data[id].Code
	}

	for i, size := range hashMapSizes {
		hashMap = postnum.CreateHashMap(data, size)
		times := make([]int64, repeats)
		bucketSizes := make([]int64, repeats)

		for i := 0; i < repeats; i++ {
			start := time.Now()
			_, bucketSize := hashMap.Lookup(lookupZipCodes[i])
			time1 := time.Since(start).Nanoseconds()

			bucketSizes[i] = bucketSize
			times[i] = time1
		}

		hashMapTimesVaried[i] = times
		hashMapBucketSizes[i] = bucketSizes
	}

	hashMapSizesHeader := make([]string, len(hashMapSizes))

	for i, size := range hashMapSizes {
		hashMapSizesHeader[i] = fmt.Sprintf("%d", size)
	}

	WriteTimes("hashmap_times.csv", hashMapTimesVaried, hashMapSizesHeader)
	WriteTimes("hashmap_bucket_sizes.csv", hashMapBucketSizes, hashMapSizesHeader)
}
