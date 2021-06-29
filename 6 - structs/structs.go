package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	lograduro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)
	u.nome = "erick"
	u.idade = 27
	fmt.Println(u)

	endereco1 := endereco{"Rua dos bobos", 8}

	// inferencia
	usuario2 := usuario{"erick", 27, endereco1}
	fmt.Println(usuario2)

	// Inicializar sin valor de otros campos
	usuario3 := usuario{nome: "erick"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 28}
	fmt.Println(usuario4)

}