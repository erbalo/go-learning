package main

import "fmt"

var n int

func init() { // siempre se ejecuta antes que el main, una por archivo no por paquete
	fmt.Println("Executando a funcao init")
	n = 10
}

func main() {
	fmt.Println("Funcao main sendo executada")
	fmt.Println(n)
}
