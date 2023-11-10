package algoritmos

import (
	"fmt"
	"proyecto_final_algoritmos/matrices"
	"time"
)

func LlamarFloydWarshall(tam int) int {
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

	floydWarshall(graph, inicio, fin)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
	return int(elapsedTime)

}

func floydWarshall(dist [][]int, src int, dest int) int {
	V := len(dist)

	for k := 0; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist[src][dest]
}
