package algoritmos

import (
	"fmt"
	"proyecto_final_algoritmos/matrices"
	"sync"
	"time"
)

func LlamarDijkstraParalelo(tam int) int{
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


	parallelDijkstra(graph, inicio, fin)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
	return int(elapsedTime)
}

// Algoritmo de Dijkstra en paralelo para encontrar la distancia más corta entre dos nodos
func parallelDijkstra(adj [][]int, src, dst int) int {
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
	for len(open) > 0 {
		// Extraemos el nodo con la distancia más baja de la cola
		mu.Lock()
		nodo := <-open
		mu.Unlock()

		// Para cada vecino del nodo
		for vecino := range adj[nodo] {
			// Si el vecino no ha sido explorado y la distancia al vecino a través del nodo es menor que la distancia actual al vecino:
			if distances[vecino] == int(^uint(0)>>1) && distances[nodo]+adj[nodo][vecino] < distances[vecino] {
				// Actualizamos la distancia al vecino
				distances[vecino] = distances[nodo] + adj[nodo][vecino]

				// Añadimos el vecino a la cola
				open <- vecino
			}
		}
	}

	// Devolvemos la distancia entre el nodo origen y el destino
	return distances[dst]
}
