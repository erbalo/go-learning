package main

import (
	"errors"
	"fmt"
)

func main() {
	// int(arquitectura de la computadora), int8, int16, int32, int64
	var numero int = 10000000000
	fmt.Println(numero)

	// uint unsigned int
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.54
	fmt.Println(numeroReal3)

	var str string = "texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	// no existen char, siempre se tomaran como int por la tabla ascii
	char := 'B'
	fmt.Println(char)

	// valor zero, todo tipo de dato giene el valor inicial string (vacio) int (0) etc ...
	var texto float32
	fmt.Println(texto)

	var boolean1 bool = true
	fmt.Println(boolean1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}