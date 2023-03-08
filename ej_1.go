//1)	Haga un programa que cuente e indique el número de caracteres, el número de palabras y el número de líneas
//de un texto cualquiera (obtenido de cualquier forma que considere). Considere hacer el cálculo de las tres
//variables en el mismo proceso.

package main

import (
	"fmt"
	"strings"
)

func main() {

	txt := "Mi nombre es Luxgie Murillo"

	// Contar caracteres
	num_caracteres := len(txt)

	// Contar palabras
	palabras := strings.Fields(txt)
	num_palabras := len(palabras)

	// Contar líneas
	num_lineas := 1
	for _, char := range txt {
		if char == '\n' {
			num_lineas++
		}
	}

	// Imprimir resultados
	fmt.Println("Caracteres:", num_caracteres)
	fmt.Println("Palabras:", num_palabras)
	fmt.Println("líneas:", num_lineas)
}
