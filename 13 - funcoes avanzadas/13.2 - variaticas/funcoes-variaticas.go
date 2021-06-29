package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

// solo se puede tener un parametro variatico por función y debe ser el ultimo argumento de la funcion
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funcoes variaticas")

	total := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	escrever("ola mundo", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
