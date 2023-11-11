package algoritmos

import (
	"fmt"
	"math"
	"proyecto_final_algoritmos/matrices"
	"time"
)

func LlamarGoldberg(tam int) float64 {
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

	GoldbergTarjan(graph, inicio, fin)
	elapsedTime := time.Since(startTime).Seconds() // Convierte la duración a segundos
	fmt.Printf("Tiempo de ejecución: %.6f segundos\n", elapsedTime)
	return elapsedTime
}

const infinity = math.MaxInt32

// GoldbergTarjan encuentra el camino más corto entre dos nodos en un grafo representado como una matriz de adyacencia.
func GoldbergTarjan(graph [][]int, source, destination int) ([]int, int) {
	n := len(graph)
	distance := make([]int, n)
	parent := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		distance[i] = infinity
		parent[i] = -1
		visited[i] = false
	}

	distance[source] = 0

	for i := 0; i < n-1; i++ {
		u := min(distance, visited)
		visited[u] = true

		for v := 0; v < n; v++ {
			if !visited[v] && graph[u][v] != 0 && distance[u]+graph[u][v] < distance[v] {
				distance[v] = distance[u] + graph[u][v]
				parent[v] = u
			}
		}
	}

	// Reconstruir el camino más corto
	path := []int{}
	current := destination
	for current != -1 {
		path = append([]int{current}, path...)
		current = parent[current]
	}

	return path, distance[destination]
}

// min encuentra el nodo no visitado con la distancia mínima
func min(distance []int, visited []bool) int {
	min := infinity
	minIndex := -1

	for i, d := range distance {
		if !visited[i] && d <= min {
			min = d
			minIndex = i
		}
	}

	return minIndex
}
