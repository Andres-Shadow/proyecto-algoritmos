package algoritmos

import (
	"fmt"
	"math"
	"proyecto_final_algoritmos/matrices"
	"time"
)

func LlamarYen(tam int) int{
	startTime := time.Now()
	var filas int
	var nombre string

	switch tam {
	case 1:
		filas = 1024
		nombre = "./datos/Matriz1.dat"
		break
	case 2:
		filas = 2048
		nombre = "./datos/Matriz2.dat"
		break
	case 3:
		filas = 4096
		nombre = "./datos/Matriz3.dat"
		break
	case 4:
		filas = 8192
		nombre = "./datos/Matriz4.dat"
	}

	graph, _ := matrices.LoadMatrixFromFile(nombre, filas, filas)

	inicio := 0
	fin := len(graph) - 1

	yen(graph, inicio, fin)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
	return int(elapsedTime)
}

func yen(graph [][]int, origin, destination int) []int {
	distances := make([]float64, len(graph))
	predecessors := make([]int, len(graph))

	for i := range distances {
		distances[i] = math.Inf(1)
		predecessors[i] = -1
	}
	distances[origin] = 0

	for range graph {
		for v := range graph {
			for u := range graph {
				if graph[v][u] != 0 && distances[u] > distances[v]+float64(graph[v][u]) {
					distances[u] = distances[v] + float64(graph[v][u])
					predecessors[u] = v
				}
			}
		}
	}

	shortestPath := []int{destination}
	pred := predecessors[destination]

	for pred != -1 && pred != origin {
		shortestPath = append([]int{pred}, shortestPath...) // Corregido aquí
		pred = predecessors[pred]
	}

	if pred != origin {
		return nil
	}

	shortestPath = append([]int{origin}, shortestPath...) // Agregar el origen al inicio del camino
	return shortestPath
}
