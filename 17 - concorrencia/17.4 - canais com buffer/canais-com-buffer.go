package main

import "fmt"

func main() {
	canal := make(chan string, 2) // segundo parametro es la capacidad del buffer para el numero de oepraciones en este caso 2
	// bloquea por que no tiene quien reciba mas adelante el valor, causara un deadlock
	canal <- "Ola mundo"
	canal <- "Prgroamando em go"
	//canal <- "Tercer valor" si se descomenta esta linea, como el buffer es de 2 causara un deadlock, porque supera la capacidad del canal

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
