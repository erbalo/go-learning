package main

import "fmt"

func main() {
	soma := 1 + 2
	substracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 3
	restoDivisao := 10 % 3

	fmt.Println(soma, substracao, divisao, multiplicacao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25

	// no se pueden sumar o operar tipos diferentes
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	// operadores de atribucion
	var variavel1 string = "erick"
	variavel2 := "erick2"
	fmt.Println(variavel1, variavel2)

	// relacionales
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// operadores logicos
	fmt.Println("--------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// operadores unarios
	numero := 10
	numero = numero + 1
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero += 10
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	// operador ternario, no existe
	var texto string
	if numero > 5 { // hack
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}