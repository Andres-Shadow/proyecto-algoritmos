package hilos

import (
	"proyecto_final_algoritmos/algoritmos"
	"proyecto_final_algoritmos/matrices"
	"strconv"
)

/*
func LlamarAlgoritmosConcurrentes(tam int) {
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex para proteger el acceso al archivo

	var nombre = "./resultados/resultados1.txt"

	algoritmoConcurrente := func(nombreAlgoritmo string, funcionAlgoritmo func(int) int) {
		defer wg.Done()

		a := funcionAlgoritmo(tam)

		texto := nombreAlgoritmo + ", " + strconv.Itoa(a)

		// Proteger el acceso al archivo usando el mutex
		mu.Lock()
		matrices.GuardarEnArchivo(texto, nombre)
		mu.Unlock()
	}

	// Lista de algoritmos a ejecutar concurrentemente
	algoritmos := map[string]func(int) int{
		"Spfa":                  algoritmos.LlamarSpfa,
		"Yen":                   algoritmos.LlamarYen,
		"BellmanFord":           algoritmos.LlamarBellmanFord,
		"AstarParalelo":         algoritmos.LlamarAStarParalelo,
		"FloydWarshall":         algoritmos.LlamarFloydWarshall,
		"Dijkstra":              algoritmos.LlamarDijkstra,
		"Bfs":                   algoritmos.LlamarBfs,
		"DijkstraParalelo":      algoritmos.LlamarDijkstraParalelo,
		"Ucs":                   algoritmos.LlamarUcs,
		"FloydWarshallParalelo": algoritmos.LlamarFloydWarshallParalelo,
		"AStar":                 algoritmos.LlamarAStar,
		"Goldberg":              algoritmos.LlamarGoldberg,
	}

	// Iniciar las goroutines
	for nombreAlgoritmo, funcionAlgoritmo := range algoritmos {
		wg.Add(1)
		go algoritmoConcurrente(nombreAlgoritmo, funcionAlgoritmo)
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
}*/

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

	var a int
	var texto string
	a = algoritmos.LlamarSpfa(tam)
	texto = "Spfa, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarYen(tam)
	texto = "Yen, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarBellmanFord(tam)
	texto = "BellmanFord, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarAStarParalelo(tam)
	texto = "Astar Paralelo, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarFloydWarshall(tam)
	texto = "FloydWarshall, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarDijkstra(tam)
	texto = "Dijkstra, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarBfs(tam)
	texto = "Bfs, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarDijkstraParalelo(tam)
	texto = "Dijkstra Paralelo, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarUcs(tam)
	texto = "Ucs, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarFloydWarshallParalelo(tam)
	texto = "FloydWarshall Paralelo, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarAStar(tam)
	texto = "AStar, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)
	a = algoritmos.LlamarGoldberg(tam)
	texto = "Goldberg, " + strconv.Itoa(a)
	matrices.GuardarEnArchivo(texto, nombre)

}
