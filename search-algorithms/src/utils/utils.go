package utils

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/savio04/special-topics-in-computing/search-algotithms/algorithms/binarysearch"
	"github.com/savio04/special-topics-in-computing/search-algotithms/algorithms/cubicsearch"
	"github.com/savio04/special-topics-in-computing/search-algotithms/algorithms/linearsearch"
	"github.com/savio04/special-topics-in-computing/search-algotithms/algorithms/quadraticsearch"
	"github.com/savio04/special-topics-in-computing/search-algotithms/algorithms/ternarysearch"
)

type ExecuteParam struct {
	Input             []int
	Algorithm         string
	NumberRepetitions int
	Search            int
}

type ReturnExecuteParams struct {
	Casesize           int
	Algorithm          string
	Totalexecutiontime time.Duration
	Averagerunningtime time.Duration
	TotalmemoryUsed    uint64
	AvaragememoryUsed  uint64
}

func Runalgorithm(params ExecuteParam) ReturnExecuteParams {
	runtime.GC()
	var totalTime = time.Duration(0)
	var totalMemory uint64

	var memStatsBefore, memStatsAfter runtime.MemStats

	runtime.ReadMemStats(&memStatsBefore)

	start := time.Now()

	for index := 0; index < params.NumberRepetitions; index++ {
		var result int
		switch params.Algorithm {
		case "linearsearchv1":
			result = linearsearch.Linearsearchv1(params.Input, params.Search)
		case "linearsearchv2":
			result = linearsearch.Linearsearchv2(params.Input, params.Search)
		case "binarysearchstart":
			result = binarysearch.Binarysearch(params.Input, params.Search, 0, len(params.Input)-1)
		case "binarysearchmiddle":
			result = binarysearch.Binarysearch(params.Input, params.Search, 0, len(params.Input)-1)
		case "quadraticsearch":
			result = quadraticsearch.Quadraticsearch(params.Input, params.Search)
		case "ternarysearch":
			result = ternarysearch.Ternarysearch(params.Input, params.Search)
		case "cubicsearch":
			result = cubicsearch.Cubicsearch(params.Input, params.Search)
		default:
			fmt.Println("Algorithm does not exist")
		}

		fmt.Println(result)
	}

	totalTime = time.Since(start)

	runtime.ReadMemStats(&memStatsAfter)

	totalMemory = memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc

	returnProps := ReturnExecuteParams{
		Casesize:           len(params.Input),
		Algorithm:          params.Algorithm,
		Totalexecutiontime: totalTime,
		Averagerunningtime: totalTime / time.Duration(params.NumberRepetitions),
		TotalmemoryUsed:    totalMemory,
		AvaragememoryUsed:  totalMemory / uint64(params.NumberRepetitions),
	}

	return returnProps
}

func Mapstringarraytointarray(stringArray []string) ([]int, error) {
	intArray := make([]int, len(stringArray))

	for i, str := range stringArray {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		intArray[i] = num
	}

	return intArray, nil
}

func Readfile(path string) ([]int, error) {
	content, error := ioutil.ReadFile(path)

	if error != nil {
		log.Fatal(error)
	}

	str := string(content)

	lines := strings.Split(str, "\n")
	array := make([]string, 0)

	for _, eachline := range lines {
		numStrings := strings.Split(eachline, ",")

		if len(numStrings) > 1 {
			array = append(array, numStrings...)
		}
	}

	return Mapstringarraytointarray(array)
}

func Concatenatepaths(filenames []string, inputpath string) []string {
	fullpaths := make([]string, 0)

	for _, eachfilename := range filenames {
		fullpaths = append(fullpaths, inputpath+"/"+eachfilename)
	}

	return fullpaths
}

func WriteCSV(filename string, data [][]string) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writerCSV := csv.NewWriter(file)
	defer writerCSV.Flush()

	for _, line := range data {
		if err := writerCSV.Write(line); err != nil {
			log.Fatal(err)
		}
	}
}
