package matrices

import (
	"fmt"
	"os"
)

func GuardarEnArchivo(texto string, nombreArchivo string) error {
	// Verifica si el archivo ya existe
	//a
	_, err := os.Stat(nombreArchivo)

	// Si el archivo existe, abre el archivo en modo append
	if err == nil {
		file, err := os.OpenFile(nombreArchivo, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		// Escribe el texto al final del archivo
		_, err = file.WriteString(texto+"\n")
		if err != nil {
			return err
		}
	} else {
		// Si el archivo no existe, créalo y escribe el texto
		file, err := os.Create(nombreArchivo)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(texto+"\n")
		if err != nil {
			return err
		}
	}

	fmt.Printf("Texto guardado en el archivo %s\n", nombreArchivo)
	return nil
}
