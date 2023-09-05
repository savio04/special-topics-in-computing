package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type DataGraph struct {
	times    []string
	memories []float64
}

func main() {
	algorithms := []string{
		"linearsearchv1",
		"linearsearchv2",
		"binarysearchstart",
		"binarysearchmiddle",
		"ternarysearch",
		"quadraticsearch",
		"cubicsearch",
	}

	headers := []string{"Busca Linear v1", "Busca Linear v2", "Busca Binaria (primeiro)", "Busca Binaria (meio)", "Busca Quadratica", "Busca Cubica"}
	data := make(map[string]DataGraph)

	for _, algorithm := range algorithms {
		path := "results/performance/" + algorithm + ".csv"
		content, error := ioutil.ReadFile(path)

		if error != nil {
			log.Fatal(error)
		}

		str := string(content)

		lines := strings.Split(str, "\n")
		lines = lines[1:]

		for _, line := range lines {

			items := strings.Split(line, ",")

			if !(len(items) > 1) {
				continue
			}

			size := items[len(items)-1]
			time := items[1]
			memoryString := items[2]
			memory, _ := strconv.ParseFloat(strings.Split(memoryString, " ")[0], 64)

			if _, ok := data[size]; !ok {
				data[size] = DataGraph{
					times:    make([]string, 0), // Inicialize as slices vazias
					memories: make([]float64, 0),
				}
			} else {
				dataGraph := data[size]

				dataGraph.times = append(dataGraph.times, time)
				dataGraph.memories = append(dataGraph.memories, memory)

				data[size] = dataGraph
			}
		}
	}

	for size, item := range data {
		// ptime := plot.New()
		pmemory := plot.New()

		// ptime.Title.Text = "Comparação de Desempenho para Tamanho de Entrada " + size
		// ptime.X.Label.Text = "Algoritmo"
		// ptime.Y.Label.Text = "Tempo Médio (µs)"

		pmemory.Title.Text = "Comparacao de Desempenho para Tamanho de Entrada" + size
		pmemory.X.Label.Text = "Algoritmo"
		pmemory.Y.Label.Text = "Memoria media (bytes)"

		bars, err := plotter.NewBarChart(plotter.Values(item.memories), vg.Points(40))

		if err != nil {
			log.Fatal(err)
		}

		bars.LineStyle.Width = vg.Length(0)
		bars.Color = plotutil.Color(2)

		pmemory.Add(bars)

		pmemory.NominalX(headers...)

		filenameForMemory := "results/graphics/" + "performance_comparison_memory_" + size + ".pdf"

		if err := pmemory.Save(8*vg.Inch, 4*vg.Inch, filenameForMemory); err != nil {
			log.Fatal(err)
		}
	}
}
