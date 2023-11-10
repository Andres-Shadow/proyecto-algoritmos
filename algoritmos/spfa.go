package algoritmos

import (
	"fmt"
	"math"
	"proyecto_final_algoritmos/matrices"
	"time"
)

func LlamarSpfa(tam int) int{
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

	spfa(graph, inicio, fin)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
	return int(elapsedTime)
}

func spfa(graph [][]int, origin, destination int) []int {
	distances := make([]float64, len(graph))
	for i := range distances {
		distances[i] = math.Inf(1)
	}
	distances[origin] = 0

	queue := []int{origin}

	// Map para rastrear los predecesores en el camino más corto
	predecessors := make(map[int]int)
	predecessors[origin] = -1 // El nodo de origen no tiene predecesor

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for neighbor, weight := range graph[node] {
			if weight > 0 && distances[node]+float64(weight) < distances[neighbor] {
				distances[neighbor] = distances[node] + float64(weight)
				predecessors[neighbor] = node
				queue = append(queue, neighbor)
			}
		}
	}

	// Reconstruir el camino más corto desde el nodo de destino hasta el origen
	shortestPath := []int{destination}
	for pred := predecessors[destination]; pred != -1; pred = predecessors[pred] {
		shortestPath = append([]int{pred}, shortestPath...)
	}

	if len(shortestPath) == 1 && shortestPath[0] != origin {
		// Si no hay camino o el camino encontrado no comienza en el nodo de origen
		return nil
	}

	return shortestPath
}
