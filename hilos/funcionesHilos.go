package hilos

import (
	"proyecto_final_algoritmos/algoritmos"
	"proyecto_final_algoritmos/matrices"
	"strconv"
)
func LlamarAlgoritmos(tam int) {

	var nombre = "./resultados/resultados1.txt"

	switch tam {
	case 2:
		nombre = "./resultados/resultados2.txt"
		break
	case 3:
		nombre = "./resultados/resultados3.txt"
		break
	case 4:
		nombre = "./resultados/resultados4.txt"
		break
	}

	var a float64
	var texto string

	matrices.EliminarArchivo(nombre)

	a = algoritmos.LlamarSpfa(tam)
	texto = "Spfa, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarYen(tam)
	texto = "Yen, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarBellmanFord(tam)
	texto = "BellmanFord, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarAStarParalelo(tam)
	texto = "Astar Paralelo, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarFloydWarshall(tam)
	texto = "FloydWarshall, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarDijkstra(tam)
	texto = "Dijkstra, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarBfs(tam)
	texto = "Bfs, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarDijkstraParalelo(tam)
	texto = "Dijkstra Paralelo, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarUcs(tam)
	texto = "Ucs, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarFloydWarshallParalelo(tam)
	texto = "FloydWarshall Paralelo, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarAStar(tam)
	texto = "AStar, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)

	a = algoritmos.LlamarGoldberg(tam)
	texto = "Goldberg, " + strconv.FormatFloat(a, 'f', -1, 32)
	matrices.GuardarEnArchivo(texto, nombre)
}
