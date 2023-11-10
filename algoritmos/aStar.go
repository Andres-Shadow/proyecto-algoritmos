package algoritmos

import (
	"container/heap"
	"fmt"
	"proyecto_final_algoritmos/matrices"
	"time"
)

func LlamarAStar(tam int) int{
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

	aEstrellaIterativo(graph, inicio, fin)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
	return int(elapsedTime)
}

type elemento struct {
	node      int
	prioridad int
}

type ColaDePrioridad []*elemento

func (cp ColaDePrioridad) Len() int { return len(cp) }

func (cp ColaDePrioridad) Less(i, j int) bool {
	return cp[i].prioridad < cp[j].prioridad
}

func (cp ColaDePrioridad) Swap(i, j int) {
	cp[i], cp[j] = cp[j], cp[i]
}

func (cp *ColaDePrioridad) Push(x interface{}) {
	item := x.(*elemento)
	*cp = append(*cp, item)
}

func (cp *ColaDePrioridad) Pop() interface{} {
	old := *cp
	n := len(old)
	item := old[n-1]
	*cp = old[0 : n-1]
	return item
}

func heuristica(nodo, fin int) int {
	return abs(nodo - fin)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func aEstrellaIterativo(grafico [][]int, inicio, fin int) ([]int, int) {
	colaPrioridad := make(ColaDePrioridad, 0)
	heap.Init(&colaPrioridad)
	heap.Push(&colaPrioridad, &elemento{node: inicio, prioridad: 0})
	vinoDe := make(map[int]int)
	puntuacionG := make([]int, len(grafico))
	for i := range puntuacionG {
		puntuacionG[i] = int(^uint(0) >> 1) // Inicializar con "infinito"
	}
	puntuacionG[inicio] = 0

	for colaPrioridad.Len() > 0 {
		item := heap.Pop(&colaPrioridad).(*elemento)
		nodoActual := item.node

		if nodoActual == fin {
			ruta := make([]int, 0)
			for {
				ruta = append(ruta, nodoActual)
				previo, existe := vinoDe[nodoActual]
				if !existe {
					break
				}
				nodoActual = previo
			}
			rutaInvertida := make([]int, len(ruta))
			for i, j := 0, len(ruta)-1; j >= 0; i, j = i+1, j-1 {
				rutaInvertida[i] = ruta[j]
			}
			return rutaInvertida, puntuacionG[fin]
		}

		for vecino, costo := range grafico[nodoActual] {
			if costo > 0 {
				puntuacionGtentativa := puntuacionG[nodoActual] + costo
				if puntuacionGtentativa < puntuacionG[vecino] {
					vinoDe[vecino] = nodoActual
					puntuacionG[vecino] = puntuacionGtentativa
					puntuacionF := puntuacionGtentativa + heuristica(vecino, fin)
					heap.Push(&colaPrioridad, &elemento{node: vecino, prioridad: puntuacionF})
				}
			}
		}
	}

	return nil, int(^uint(0) >> 1)
}
