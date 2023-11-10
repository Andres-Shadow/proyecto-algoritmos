package main

import (
	"fmt"
	"os"
	"proyecto_final_algoritmos/algoritmos"
	"proyecto_final_algoritmos/hilos"
	"proyecto_final_algoritmos/matrices"
	"strconv"
)

func main() {
	opcion := os.Args[1]
	var tam int

	if len(os.Args) < 3 {
		fmt.Println("1. primer tamaño")
		fmt.Println("2. segundo tamaño")
		fmt.Println("3. tercer tamaño")
		fmt.Println("4. cuarto tamaño")
		fmt.Print("Seleccione el tamaño de la matriz: ")
		fmt.Scan(&tam)
	} else {
		tam, _ = strconv.Atoi(os.Args[2])
	}

	switch opcion {
	case "1":
		algoritmos.LlamarSpfa(tam)
		break
	case "2":
		algoritmos.LlamarYen(tam)
		break
	case "3":
		algoritmos.LlamarBellmanFord(tam)
		break
	case "4":
		algoritmos.LlamarAStarParalelo(tam)
		break
	case "5":
		algoritmos.LlamarFloydWarshall(tam)
		break
	case "6":
		algoritmos.LlamarDijkstra(tam)
		break
	case "7":
		algoritmos.LlamarBfs(tam)
		break
	case "8":
		algoritmos.LlamarDijkstraParalelo(tam)
		break
	case "9":
		algoritmos.LlamarUcs(tam)
		break
	case "10":
		algoritmos.LlamarFloydWarshallParalelo(tam)
		break
	case "11":
		algoritmos.LlamarAStar(tam)
		break
	case "12":
		algoritmos.LlamarGoldberg(tam)
		break
	case "x":
		matrices.GenearMatricesPorTamanio(tam)
		break
	case "a":
		 hilos.LlamarAlgoritmos(1)
		 hilos.LlamarAlgoritmos(2)
		 hilos.LlamarAlgoritmos(3)
		 hilos.LlamarAlgoritmos(4)
	}
}
