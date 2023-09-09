package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/phanty133/id1021/stack/pkg/solver"
	"github.com/phanty133/id1021/stack/pkg/stacks"
)

func StackBench[StackType stacks.Stack[int]](tag string, stack StackType, runs int, runIters int, stackIters int) {
	outFile, err := os.Create(fmt.Sprintf("stack_%s.csv", tag))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer outFile.Close()

	runTimes := make([]time.Duration, runs)

	for run := 0; run < runs; run++ {
		runStart := time.Now()

		for iter := 0; iter < runIters; iter++ {
			for i := 0; i < stackIters; i++ {
				stack.Push(i)
			}
	
			for i := 0; i < stackIters; i++ {
				stack.Pop()
			}
		}

		runTimes[run] = time.Since(runStart)
	}

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	for _, time := range runTimes {
		writer.Write([]string{fmt.Sprintf("%d", time.Microseconds())})
	}
}

func Bench() {
	runs := 2000
	iters := 1000
	stackOps := [5]int{100, 500, 1000, 2000, 5000}
	
	for _, op := range stackOps {
		staticData := make([]int, op * 2)
		staticStack := stacks.NewStaticStack[int](staticData)
		dynStack := stacks.NewDynamicStack[int](4)

		StackBench(fmt.Sprintf("dynamic-%d", op), dynStack, runs, iters, op)
		StackBench(fmt.Sprintf("static-%d", op), staticStack, runs, iters, op)
	}
}

func main() {
	// Bench()

	// dataArr := [10]float32{}
	// numStack := stacks.NewStaticStack[float32](dataArr[:])

	numStack := stacks.NewDynamicStack[float32](10)
	expr := "0.5 4 8 * +"

	result, err := solver.Solve(numStack, expr)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
