package main

import (
	"fmt"
	"time"
)

func main() {
	// canal tanto enviar como recibir datos
	// El canal tiene dos operaciones, recibir y enviar datos, bloquean la ejecucion del programa
	canal := make(chan string) // traficar datos entre el canal, en esta caso string
	go escrever("Ola Mundo!", canal)

	fmt.Println("depois da funcao escrever comezar a ser executada")

	// deadLock esta esperando a recibir un dato infinitamente pero como no llego nada causa un deadlock, solo ocurre en ejecuación y no en compilación
	/*for {
		mensagem, aberto := <-canal // espera recibir el valor, la segunda varaible es saber si esta abierto o no, son operaciones bloqueantes

		if !aberto {
			break
		}

		fmt.Println(mensagem)
	}*/

	// otra forma de hacer lo mismo de arriba
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fin do programa")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- fmt.Sprintf("%s - %d", texto, i) // el texto se está enviando al canal
		time.Sleep(time.Second)
	}

	close(canal)
}
