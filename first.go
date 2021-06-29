package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ola Mundo!")
	// go mod init <module_name> inicia un modulo para obtener dependencias
	// go build crea el ejecutable
	// go install es casi lo mismo que go build solo que instala en el root de instalacion de go
	// go mod tidy --> quita las dependencias que no son usadas
}