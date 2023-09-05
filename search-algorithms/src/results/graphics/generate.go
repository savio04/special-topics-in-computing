package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	algoritmos := []string{"Busca Linear v1", "Busca Linear v2", "Busca Binária (primeiro)", "Busca Binária (meio)", "Busca Quadrática", "Busca Cúbica"}
	temposMedios := []float64{3.424, 5.23, 3.887, 4.29, 7.86, 761.346}

	p := plot.New()

	p.Title.Text = "Comparação de Desempenho para Tamanho de Entrada 100"
	p.X.Label.Text = "Algoritmo"
	p.Y.Label.Text = "Tempo Médio (µs)"

	bars, err := plotter.NewBarChart(plotter.Values(temposMedios), vg.Points(40))
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(0)

	p.Add(bars)

	p.NominalX(algoritmos...)

	if err := p.Save(8*vg.Inch, 4*vg.Inch, "results/graphics/performance_comparison.pdf"); err != nil {
		log.Fatal(err)
	}
}
