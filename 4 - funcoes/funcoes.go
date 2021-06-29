package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// si no se escpefica el tipo toma el ultimo declarado
func calcuosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	substracao := n1 - n2
	return soma, substracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Print("função f ")
		fmt.Println(txt)

		return txt
	}

	resultado := f("Texto na função")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubstracao := calcuosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubstracao)

	// solo tomamos el primer resultado
	resultadoSoma2, _ := calcuosMatematicos(10, 10)
	fmt.Println(resultadoSoma2)

	// solo tomamos el segundo resultado
	_ , resultadoSubstracao2 := calcuosMatematicos(8,2)
	fmt.Println(resultadoSubstracao2)
}