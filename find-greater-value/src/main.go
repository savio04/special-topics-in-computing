package main

import (
	"fmt"
	"log"

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

	inputPaths := utils.Concatenatepaths(fileNames, basePathNotordained)

	for _, eachFilePath := range inputPaths {
		cases, err := utils.Readfile(eachFilePath)

		if err != nil {
			log.Fatal(err)
		}

		resultv2 := maximumvalue.Maximumvaluev2(cases, 0, len(cases)-1)

		fmt.Println(resultv2)

		fmt.Println()
	}
}
