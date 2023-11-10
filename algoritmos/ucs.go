package algoritmos

import (
	"container/heap"
	"fmt"
	"proyecto_final_algoritmos/matrices"
	"time"
)

func LlamarUcs(tam int)int {
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

	ucs(graph, inicio, fin)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
	return int(elapsedTime)
}

type Item struct {
	cost  int
	node  int
	path  []int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func ucs(graph [][]int, start, end int) []int {
	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)
	item := &Item{cost: 0, node: start, path: []int{}}
	heap.Push(&priorityQueue, item)
	visited := make(map[int]bool)

	for priorityQueue.Len() > 0 {
		item := heap.Pop(&priorityQueue).(*Item)
		cost, node, path := item.cost, item.node, item.path

		if _, exists := visited[node]; !exists {
			path = append(path, node)
			if node == end {
				return path
			}
			visited[node] = true
			for neighbor, weight := range graph[node] {
				if weight > 0 && !visited[neighbor] {
					newCost := cost + weight
					newItem := &Item{cost: newCost, node: neighbor, path: append([]int(nil), path...)}
					heap.Push(&priorityQueue, newItem)
				}
			}
		}
	}

	return nil
}
