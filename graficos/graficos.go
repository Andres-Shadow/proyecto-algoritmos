package graficos

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type DataPoint struct {
	Name  string
	Value float64
}

func Graficar(tam int) {
	var rutaResultados string
	switch tam{
	case 1:
		rutaResultados = "./resultados/resultados1.txt"
		break
	case 2:
		rutaResultados = "./resultados/resultados2.txt"
		break
	case 3:
		rutaResultados = "./resultados/resultados3.txt"
		break
	case 4:
		rutaResultados = "./resultados/resultados4.txt"
		break
	}
	p := plot.New()

	p.Title.Text = "Rendimiento Algoritmos"
	p.X.Label.Text = "Algoritmos"
	p.Y.Label.Text = "Tiempo en Segundos"

	// Leer datos desde un archivo CSV
	data, err := readCSV(rutaResultados)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Crear un gráfico de barras
	bars, _ := plotter.NewBarChart(plotter.Values(dataValues(data)), 0.5)

	bars.LineStyle.Width = vg.Length(3)
	bars.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	// Agregar el gráfico de barras al gráfico
	p.Add(bars)

	// Agregar una línea en el eje X
	p.Add(plotter.NewGrid())

	// Establecer las etiquetas del eje X
	p.NominalX(dataLabels(data)...)
	p.X.Tick.Label.Rotation = 45 // Rotar las etiquetas del eje X para mayor claridad

	p.X.Tick.Label.XAlign = draw.XRight
	p.X.Tick.Label.YAlign = draw.YTop

	// Ajustar la escala del eje Y
	minY := 0.0
	maxY := calculateYMax(data) // Puedes definir tu propia función para calcular el máximo según tus necesidades
	p.Y.Min = minY
	p.Y.Max = maxY

	// Guardar el gráfico en un archivo (formato PNG en este caso)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "output_barras.png"); err != nil {
		fmt.Println(err)
		return
	}
}

// Funciones de ayuda para extraer valores y etiquetas de la matriz
func dataValues(data []DataPoint) []float64 {
	values := make([]float64, len(data))
	for i, dp := range data {
		values[i] = dp.Value
	}
	return values
}

func dataLabels(data []DataPoint) []string {
	labels := make([]string, len(data))
	for i, dp := range data {
		labels[i] = dp.Name
	}
	return labels
}

// Función para leer datos desde un archivo CSV
func readCSV(filename string) ([]DataPoint, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []DataPoint
	for _, line := range lines {
		name := line[0]
		value, err := strconv.ParseFloat(strings.TrimSpace(line[1]), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, DataPoint{Name: name, Value: value})
	}

	return data, nil
}

// Función para calcular el máximo valor en el eje Y
func calculateYMax(data []DataPoint) float64 {
	max := 0.0
	for _, dp := range data {
		if dp.Value > max {
			max = dp.Value
		}
	}
	return max
}
