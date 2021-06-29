package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, substracao int) {
	soma = n1 + n2 // usa retorno nombrado
	substracao = n1 - n2
	return
}

func main() {
	fmt.Println("Funcoes avanzadas")

	soma, substracao := calculosMatematicos(10, 5)
	fmt.Println(soma, substracao)
}
