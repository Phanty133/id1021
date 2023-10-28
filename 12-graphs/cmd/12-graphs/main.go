package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/phanty133/id1021/12-graphs/pkg/trains"
	"os"
	"strconv"
	"sync"
	"time"
)

var DATA_PATH string = "/home/phanty/repos/id1021/12-graphs/data/trains.csv"

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

func ReadData() ([]string, []string, []int) {
	from := make([]string, 0)
	to := make([]string, 0)
	dist := make([]int, 0)

	file, err := os.Open(DATA_PATH)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ','

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		from = append(from, record[0])
		to = append(to, record[1])

		distInt, _ := strconv.Atoi(record[2])
		dist = append(dist, distInt)
	}

	return from, to, dist
}

func BenchNaive(graph trains.NetworkGraph, pairs [][]string) {
	var wg sync.WaitGroup

	for _, pair := range pairs {
		maxDist := 300
		wg.Add(1)

		go func(a string, b string) {
			defer wg.Done()

			for {
				start := time.Now()
				d, err := graph.DepthFirstDistance(a, b, maxDist)
				elapsed := time.Since(start)

				if err == nil {
					fmt.Printf("--- Benchmarking %s -> %s\n", a, b)
					fmt.Println("Distance:", d)
					fmt.Println("Time:", elapsed)
					break
				}

				maxDist = int(float32(maxDist) * 1.5)
				// fmt.Printf("Max distance too low, increasing to %d\n", maxDist)
			}
		}(pair[0], pair[1])
	}

	wg.Wait()
}

func BenchNoLoop(graph trains.NetworkGraph, pairs [][]string) {
	var wg sync.WaitGroup

	for _, pair := range pairs {
		wg.Add(1)

		go func(a string, b string) {
			defer wg.Done()

			start := time.Now()
			d, err := graph.DepthFirstDistanceNoLoop(a, b)
			elapsed := time.Since(start)

			if err != nil {
				fmt.Printf("No path found between %s and %s\n", a, b)
				return
			}

			fmt.Printf("--- Benchmarking %s -> %s\n", a, b)
			fmt.Println("Distance:", d)
			fmt.Println("Time:", elapsed)
		}(pair[0], pair[1])
	}

	wg.Wait()
}

func BenchNoLoopMaxed(graph trains.NetworkGraph, pairs [][]string) {
	var wg sync.WaitGroup

	for _, pair := range pairs {
		wg.Add(1)

		go func(a string, b string) {
			defer wg.Done()

			start := time.Now()
			d, _, err := graph.DepthFirstDistanceNoLoopMaxed(a, b)
			elapsed := time.Since(start)

			if err != nil {
				fmt.Printf("No path found between %s and %s\n", a, b)
				return
			}

			fmt.Printf("--- Benchmarking %s -> %s\n", a, b)
			fmt.Println("Distance:", d)
			fmt.Println("Time:", elapsed)
		}(pair[0], pair[1])
	}

	wg.Wait()
}

func BenchNoLoopMaxedMalmo(graph trains.NetworkGraph) {
	FROM := "Malmö"
	REPEATS := 50

	times := make([][]int64, 0)
	header := make([]string, 0)

	for _, target := range graph.GetCities() {
		fmt.Println("Benchmarking", target)

		var d int
		var minPath []string
		repTimes := make([]int64, 0)
		header = append(header, target)

		for i := 0; i < REPEATS; i++ {
			start := time.Now()
			d, minPath, _ = graph.DepthFirstDistanceNoLoopMaxed(FROM, target)
			elapsed := time.Since(start)

			repTimes = append(repTimes, elapsed.Nanoseconds())
		}

		fmt.Println(minPath)

		repTimes = append(repTimes, int64(d))
		repTimes = append(repTimes, int64(len(minPath)))
		times = append(times, repTimes)
	}

	WriteTimes("no-loop-maxed-malmo.csv", times, header)
}

func main() {
	from, to, dist := ReadData()
	graph := trains.NewNetworkGraph()

	graph.FillData(from, to, dist)
	// collisions, bucketSizes := graph.CountBucketCollisions()

	// fmt.Println("Collisions:", collisions)
	// fmt.Println("Bucket sizes:", bucketSizes)

	// for k, v := range bucketSizes {
	// 	if v == 1 {
	// 		continue
	// 	}

	// 	fmt.Printf("Bucket %d has %d elements\n", k, v)

	// 	for _, city := range graph.Cities[k] {
	// 		fmt.Println(city.Name)
	// 	}
	// }

	BENCH_PAIRS := [][]string{
		{"Malmö", "Göteborg"},
		{"Göteborg", "Stockholm"},
		{"Malmö", "Stockholm"},
		{"Stockholm", "Sundsvall"},
		{"Stockholm", "Umeå"},
		{"Göteborg", "Sundsvall"},
		{"Sundsvall", "Umeå"},
		{"Umeå", "Göteborg"},
		{"Göteborg", "Umeå"},
		{"Malmö", "Kiruna"},
	}

	// BenchNaive(graph, BENCH_PAIRS)
	BenchNoLoop(graph, BENCH_PAIRS)
	// BenchNoLoopMaxed(graph, BENCH_PAIRS)
	// BenchNoLoopMaxedMalmo(graph)
}
