// 2) Escriba el programa más eficiente que pueda para imprimir en pantalla la figura de un rombo con ateriscos
package main

import "fmt"

func main() {
	var line_centro int
	fmt.Print("Ingrese la cantidad de asteríscos de la línea del centro, debe ser impar-positiva. ")
	fmt.Scanln(&line_centro)

	if line_centro%2 == 0 || line_centro <= 0 {
		fmt.Println("La cantidad de asteríscos debe ser impar-positiva.")
		return
	}

	imprimirRombo(line_centro)
}

//si el usuario ingresa un número par el programa terminará
func imprimirRombo(line_centro int) {
	for i := 1; i <= line_centro; i += 2 {
		for j := 1; j <= (line_centro-i)/2; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := line_centro - 2; i >= 1; i -= 2 {
		for j := 1; j <= (line_centro-i)/2; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
