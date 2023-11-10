package matrices

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GenearMatricesPorTamanio(tam int) {
	var nombre1 string
	switch tam {
	case 1:
		tam = 1024
		nombre1 = "./datos/Matriz1.dat"
		break
	case 2:
		tam = 2048
		nombre1 = "./datos/Matriz2.dat"
		break
	case 3:
		nombre1 = "./datos/Matriz3.dat"
		tam = 4096
		break
	case 4:
		tam = 8192
		nombre1 = "./datos/Matriz4.dat"
	}

	m := GenerateRandomMatrix(tam, tam)
	SaveMatrixToFile(m, nombre1)
	fmt.Println("se ha generado la matriz 1 de tamaño ", tam)
}

func GenerateRandomMatrix(rows, cols int) [][]int {
	rand.Seed(time.Now().UnixNano()) // Inicializa la semilla aleatoria con el tiempo actual

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			// Genera un número aleatorio entre 10000000 y 99999999 (8 dígitos)
			matrix[i][j] = rand.Intn(1) + 100
		}
	}
	return matrix
}

func SaveMatrixToFile(matrix [][]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range matrix {
		for _, value := range row {
			err := binary.Write(file, binary.LittleEndian, int32(value))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func LoadMatrixFromFile(filename string, rows, cols int) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			var value int32
			err := binary.Read(file, binary.LittleEndian, &value)
			if err != nil {
				return nil, err
			}
			matrix[i][j] = int(value)
		}
	}

	return matrix, nil
}
