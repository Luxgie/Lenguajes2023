/*
*
3)	Escriba una función que permita rotar una secuencia de elementos N posiciones a la izquierda o a la derecha según sea el caso en función al parámetro que se reciba. Los parámetros deben ser un puntero a un arreglo previamente creado, la cantidad de movimiento de rotación y la dirección (0 = izq y 1 = der).

A partir de esta función, escriba un programa que haga varias rotaciones cualesquiera de una secuencia de elementos previamente creada que sea inmutable. Al final debe imprimirse el resultado de cada rotación además de la secuencia original.

	Un ejemplo de rotación puede ser el siguiente:

Secuencia Original = [a, b, c, d, e, f, g, h,]
Cantidad de rotaciones = 3
Dirección = izq
Secuencia final rotada = [d, e, f, g, h, a, b, c]   Nótese que cada iteración, el elemento más a la izquierda pasó a formar parte del final de la secuencia, si la rotación fuera a la derecha, sería lo contrario
*/
package main

import "fmt"

func rotateSequence(sequence []interface{}, rotations int, direction string) []interface{} {
	length := len(sequence)
	rotated := make([]interface{}, length)

	for i := 0; i < length; i++ {
		var newIndex int

		if direction == "izq" {
			newIndex = (i + rotations) % length
		} else if direction == "der" {
			newIndex = (i + length - rotations) % length
		} else {
			return sequence
		}

		rotated[newIndex] = sequence[i]
	}

	return rotated
}

func main() {
	sequence := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h"}
	rotations := 3
	direction := "izq"

	fmt.Println("Secuencia Original:", sequence)

	for i := 1; i <= rotations; i++ {
		rotated := rotateSequence(sequence, i, direction)
		fmt.Printf("Secuencia rotada %d vez(es) hacia %s: %v\n", i, direction, rotated)
	}
}
