package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}

	}()

	for {
		select { // solo usado para concurrencia
		case mensagemCanal1 := <-canal1:
			dt1 := time.Now()
			fmt.Println(dt1.Format("2006-01-02 15:04:05.00"), mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			dt2 := time.Now()
			fmt.Println(dt2.Format("2006-01-02 15:04:05.00"), mensagemCanal2)
		}
	}
}
