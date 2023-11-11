package algoritmos

import (
	"fmt"
	"proyecto_final_algoritmos/matrices"
	"sync"
	"time"
)

func LlamarAStarParalelo(tam int) float64 {
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

	parallelAStar(graph, inicio, fin)
	elapsedTime := time.Since(startTime).Seconds() // Convierte la duración a segundos
	fmt.Printf("Tiempo de ejecución: %.6f segundos\n", elapsedTime)
	return elapsedTime
}

func parallelAStar(adj [][]int, src, dst int) int {
	// Inicializamos las distancias a todos los nodos a infinito
	distances := make([]int, len(adj))
	for i := range distances {
		distances[i] = int(^uint(0) >> 1)
	}

	// Inicializamos las distancias del nodo origen a sí mismo a 0
	distances[src] = 0

	// Creamos una cola de prioridad, que contiene los nodos que aún no han sido explorados
	open := make(chan int, len(adj))
	open <- src

	// Creamos una variable para el mutex que protege las distancias
	var mu sync.Mutex

	// Iteramos mientras la cola no esté vacía
	for {
		mu.Lock()
		if len(open) == 0 {
			mu.Unlock()
			break
		}
		node := <-open
		mu.Unlock()

		// Para cada vecino del nodo
		for vecino, peso := range adj[node] {
			// Si el vecino no ha sido explorado y la distancia al vecino a través del nodo es menor que la distancia actual al vecino:
			if distances[vecino] == int(^uint(0)>>1) || distances[node]+peso < distances[vecino] {
				// Actualizamos la distancia al vecino
				distances[vecino] = distances[node] + peso

				// Añadimos el vecino a la cola
				open <- vecino
			}
		}
	}

	close(open)

	// Verifica si la distancia al nodo destino sigue siendo infinita
	if distances[dst] == int(^uint(0)>>1) {
		// No hay camino desde el origen al destino
		return -1
	}

	// Devuelve la distancia al nodo destino
	return distances[dst]
}
