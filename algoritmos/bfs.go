package algoritmos

import (
	"container/list"
	"fmt"
	"proyecto_final_algoritmos/matrices"
	"time"
)

func LlamarBfs(tam int) int {
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

	bfs(graph, inicio, fin)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
	return int(elapsedTime)
}

func bfs(graph [][]int, start, end int) []int {
	queue := list.New()
	visited := make([]bool, len(graph))
	parents := make([]int, len(graph))

	for i := range graph {
		visited[i] = false
		parents[i] = -1
	}

	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		currentNode := queue.Remove(queue.Front()).(int)
		if currentNode == end {
			path := []int{}
			for currentNode != -1 {
				path = append(path, currentNode)
				currentNode = parents[currentNode]
			}
			reversedPath := make([]int, len(path))
			for i, j := 0, len(path)-1; j >= 0; i, j = i+1, j-1 {
				reversedPath[i] = path[j]
			}
			return reversedPath
		}

		for neighbor := 0; neighbor < len(graph[currentNode]); neighbor++ {
			if graph[currentNode][neighbor] != 0 && !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				parents[neighbor] = currentNode
			}
		}
	}

	return nil
}
