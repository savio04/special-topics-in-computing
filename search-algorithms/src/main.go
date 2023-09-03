package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/savio04/special-topics-in-computing/search-algotithms/utils"
)

type Algorithms struct {
	basepath  string
	filenames []string
}

func main() {
	basePathOrdered := "mocks/ordered"
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

	algorithms := map[string]Algorithms{
		"linearsearchv1":     {basepath: basePathNotordained, filenames: fileNames},
		"linearsearchv2":     {basepath: basePathNotordained, filenames: fileNames},
		"binarysearchstart":  {basepath: basePathOrdered, filenames: fileNames},
		"binarysearchmiddle": {basepath: basePathOrdered, filenames: fileNames},
		"ternarysearch":      {basepath: basePathOrdered, filenames: fileNames},
		"quadraticsearch":    {basepath: basePathNotordained, filenames: fileNames},
		"cubicsearch":        {basepath: basePathNotordained, filenames: fileNames},
	}

	mapAlgorithms := map[string]string{
		"linearsearchv1":     "Busca Linear v1",
		"linearsearchv2":     "Busca Linear v2",
		"binarysearchstart":  "Busca Binaria (primeiro elemento)",
		"binarysearchmiddle": "Busca Binaria (elemento do meio)",
		"quadraticsearch":    "Busca Quadratica",
		"ternarysearch":      "Busca Ternaria",
		"cubicsearch":        "Busca Cubica",
	}

	rand.Seed(time.Now().UnixNano())

	/*For in algorithms*/
	for algorithmname, data := range algorithms {
		resultByAlgorithm := []utils.ReturnExecuteParams{}
		fmt.Println("Algorithm initiated:", algorithmname)

		inputPaths := utils.Concatenatepaths(data.filenames, data.basepath)

		/*For in input files*/
		for _, eachFilePath := range inputPaths {
			cases, err := utils.Readfile(eachFilePath)

			if err != nil {
				log.Fatal(err)
			}

			var search int

			if algorithmname == "binarysearchstart" {
				search = cases[0]
			} else if algorithmname == "binarysearchmiddle" {
				middle := (len(cases) - 1) / 2

				search = cases[middle]
			} else {
				randomindex := rand.Intn(len(cases))

				search = cases[randomindex]
			}

			params := utils.ExecuteParam{
				Input:             cases,
				Algorithm:         algorithmname,
				NumberRepetitions: 4,
				Search:            search,
			}

			result := utils.Runalgorithm(params)

			resultByAlgorithm = append(resultByAlgorithm, result)
		}

		/*Save result for each algorithm*/
		filename := "results/" + algorithmname + ".csv"
		data := make([][]string, 0)

		//header
		data = append(data, []string{"Algoritmo", "Tempo medio", "Memoria media", "Tamanho da entrada"})

		for _, eachresult := range resultByAlgorithm {
			data = append(data, []string{
				mapAlgorithms[eachresult.Algorithm],
				eachresult.Averagerunningtime.String(),
				strconv.FormatUint(eachresult.AvaragememoryUsed, 10) + " bytes",
				strconv.Itoa(eachresult.Casesize),
			})
		}

		utils.WriteCSV(filename, data)
	}
}
