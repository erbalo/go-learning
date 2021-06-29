package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrencia != paralelismo
	go escrever("Ola mundo!") //goroutine
	// son funciones que no necesitan terminar para mandar a llamar la siguiente funcion
	escrever("Programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
