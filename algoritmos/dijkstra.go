package algoritmos

import (
	"fmt"
	"math"
	"proyecto_final_algoritmos/matrices"
	"time"
)

func LlamarDijkstra(tam int) float64 {
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

	dijkstra(graph, inicio, fin)
	elapsedTime := time.Since(startTime).Seconds() // Convierte la duración a segundos
	fmt.Printf("Tiempo de ejecución: %.6f segundos\n", elapsedTime)
	return elapsedTime
}

func dijkstra(graph [][]int, start, end int) {
	numNodes := len(graph)
	distances := make([]int, numNodes)
	visited := make([]bool, numNodes)
	previous := make([]int, numNodes)

	for i := range distances {
		distances[i] = int(math.Inf(1))
		visited[i] = false
		previous[i] = -1
	}

	distances[start] = 0

	for i := 0; i < numNodes-1; i++ {
		u := minDistance(distances, visited)

		visited[u] = true

		for v := 0; v < numNodes; v++ {
			if !visited[v] && graph[u][v] != 0 && distances[u]+graph[u][v] < distances[v] {
				distances[v] = distances[u] + graph[u][v]
				previous[v] = u
			}
		}
	}

	//printPath(previous, end)
}

func minDistance(distances []int, visited []bool) int {
	minDist := int(math.Inf(1))
	minIndex := -1

	for i, dist := range distances {
		if !visited[i] && dist <= minDist {
			minDist = dist
			minIndex = i
		}
	}

	return minIndex
}

func printPath(previous []int, current int) {
	if previous[current] != -1 {
		printPath(previous, previous[current])
	}

	fmt.Printf("%d ", current)
}
