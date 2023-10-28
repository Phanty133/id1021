package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/phanty133/id1021/13-dijkstra/pkg/trains"
	"os"
	"strconv"
	"time"
)

var DATA_PATH string = "/home/phanty/repos/id1021/13-dijkstra/data/trains.csv"

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

func BenchPrevBest(graph trains.NetworkGraph, pairs [][]string) {
	repeats := 100
	times := make([][]int64, 0)

	for _, pair := range pairs {
		fmt.Println("Benchmarking", pair[0], pair[1])
		repTimes := make([]int64, 0)

		for i := 0; i < repeats; i++ {
			start := time.Now()
			_, _, err := graph.DepthFirstDistanceNoLoopMaxed(pair[0], pair[1])
			elapsed := time.Since(start)

			if err != nil {
				panic(err)
			}

			repTimes = append(repTimes, elapsed.Nanoseconds())

			if i%10 == 0 {
				fmt.Println("Done ", i)
			}
		}

		times = append(times, repTimes)
	}

	header := make([]string, len(pairs))

	for i, pair := range pairs {
		header[i] = fmt.Sprintf("%s-%s", pair[0], pair[1])
	}

	WriteTimes("prevbest.csv", times, header)
}

func BenchDijkstra(graph trains.NetworkGraph, pairs [][]string) {
	repeats := 100
	times := make([][]int64, 0)

	for _, pair := range pairs {
		repTimes := make([]int64, 0)
		doneNum := 0

		for i := 0; i < repeats; i++ {
			start := time.Now()
			_, _, num, err := graph.DijkstraFind(pair[0], pair[1])
			elapsed := time.Since(start)
			doneNum = num

			if err != nil {
				panic(err)
			}

			repTimes = append(repTimes, elapsed.Nanoseconds())
		}

		repTimes = append(repTimes, int64(doneNum))
		times = append(times, repTimes)
	}

	header := make([]string, len(pairs))

	for i, pair := range pairs {
		header[i] = fmt.Sprintf("%s-%s", pair[0], pair[1])
	}

	WriteTimes("dijkstra-europe.csv", times, header)
}

func BenchDijsktraFromSingle(graph trains.NetworkGraph) {
	FROM := "Berlin"
	REPEATS := 50

	times := make([][]int64, 0)
	header := make([]string, 0)

	for _, target := range graph.GetCities() {
		fmt.Println("Benchmarking", target)

		var d int
		var doneLen int
		var minPath []string
		repTimes := make([]int64, 0)
		header = append(header, target)

		for i := 0; i < REPEATS; i++ {
			start := time.Now()
			d, minPath, doneLen, _ = graph.DijkstraFind(FROM, target)
			elapsed := time.Since(start)

			repTimes = append(repTimes, elapsed.Nanoseconds())
		}

		fmt.Println(minPath)

		repTimes = append(repTimes, int64(d))
		repTimes = append(repTimes, int64(len(minPath)))
		repTimes = append(repTimes, int64(doneLen))
		times = append(times, repTimes)
	}

	WriteTimes("dijkstra-berlin.csv", times, header)
}

func main() {
	from, to, dist := ReadData()
	graph := trains.NewNetworkGraph()
	graph.FillData(from, to, dist)

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

	BenchPrevBest(graph, BENCH_PAIRS)
	// BenchDijkstra(graph, BENCH_PAIRS)
	// BenchNoLoopMaxedMalmo(graph)

	// BenchDijsktraFromSingle(graph)
}
