package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // para sincronizar las goroutines

	waitGroup.Add(4) // cantidad de goroutines, en grupo de espera

	go func() {
		escrever("Ola mundo!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em go")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Goroutine 3")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Goroutine 4")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 1 -> 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
