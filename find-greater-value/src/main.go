package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/savio04/special-topics-in-compting/find-greater-value/algorithms/maximumvalue"
	"github.com/savio04/special-topics-in-computing/search-algotithms/utils"
)

func main() {
	basePathNotordained := "mocks/notordained"

	fileNames := []string{
		"100.txt",
		"200.txt",
		"1000.txt",
		"2000.txt",
		"5000.txt",
		"10000.txt",
		"50000.txt",
		"100000.txt",
		"500000.txt",
		"1000000.txt",
		"5000000.txt",
		"10000000.txt",
		"100000000.txt",
	}

	algorithms := []string{
		"maximumvaluev1",
		"maximumvaluev2",
	}

	inputPaths := utils.Concatenatepaths(fileNames, basePathNotordained)

	for _, eachFilePath := range inputPaths {
		runtime.GC()
		cases, err := utils.Readfile(eachFilePath)

		if err != nil {
			log.Fatal(err)
		}

		for _, eachalgorithm := range algorithms {
			var totalTime = time.Duration(0)
			var totalMemory uint64

			var memStatsBefore, memStatsAfter runtime.MemStats

			var result int

			runtime.ReadMemStats(&memStatsBefore)

			start := time.Now()

			switch eachalgorithm {
			case "maximumvaluev1":
				result = maximumvalue.Maximumvaluev1(cases, len(cases))
			case "maximumvaluev2":
				result = maximumvalue.Maximumvaluev2(cases, 0, len(cases)-1)
			}

			totalTime = time.Since(start)

			runtime.ReadMemStats(&memStatsAfter)

			totalMemory = memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc

			fmt.Println("eachalgorithm", eachalgorithm)
			fmt.Println("result", result)
			fmt.Println("totalTime", totalTime)
			fmt.Println("totalMemory", totalMemory)
			fmt.Println()
		}
	}
}
