package algoritmos

import (
	"fmt"
	"proyecto_final_algoritmos/matrices"
	"sync"
	"time"
)

func LlamarFloydWarshallParalelo(tam int)int {
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

	parallelFloydWarshall(graph, inicio, fin)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
	return int(elapsedTime)

}

func parallelFloydWarshall(adj [][]int, startNode int, endNode int) int {
	// Inicializamos la matriz de distancias a un valor grande (evita desbordamiento)
	maxValue := 999999
	distances := make([][]int, len(adj))
	for i := range distances {
		distances[i] = make([]int, len(adj))
		for j := range distances[i] {
			if i == j {
				distances[i][j] = 0
			} else if adj[i][j] != 0 {
				distances[i][j] = adj[i][j]
			} else {
				distances[i][j] = maxValue
			}
		}
	}

	// Creamos una barrera para sincronizar los gorutines
	var barrier sync.WaitGroup

	// Iteramos sobre todas las iteraciones
	for k := 0; k < len(adj); k++ {
		// Iteramos sobre todas las filas
		for i := 0; i < len(adj); i++ {
			// Iteramos sobre todas las columnas
			for j := 0; j < len(adj); j++ {
				// Si la distancia de i a j a través de k es menor que la distancia actual de i a j:
				if distances[i][j] > distances[i][k]+distances[k][j] {
					distances[i][j] = distances[i][k] + distances[k][j]
				}
			}
		}

		// Sincronizamos los gorutines
		barrier.Wait()
	}

	// Devolvemos la distancia entre el nodo de inicio y el nodo de fin
	return distances[startNode][endNode]
}
