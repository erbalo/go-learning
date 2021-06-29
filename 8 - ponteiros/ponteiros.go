package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1 // por default es una copia, estan en diferentes localidades de memoria

	fmt.Println(variavel1, variavel2)

	variavel1++ // solo afecta a la variable 1
	fmt.Println(variavel1, variavel2)

	// puntero es una referencia de memoria
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro) 

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // ambos valores son modificados por que son referencia de memoria
}