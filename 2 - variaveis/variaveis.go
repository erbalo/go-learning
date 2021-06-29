package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1" // declaracion explicita
	variavel2 := "variavel 2" // declaracion implicita (inferencia de tipo)
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala3"
		variavel4 string = "lalal4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "lalala5", "lalalala6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}