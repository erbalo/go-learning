package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa // no pasar un tipo es la simulacion de herencia, todos los campos declarados previamente se colocan en este nuevo struct
	curso string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"erick", "barrera", 27, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia", "usp"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.pessoa.nome) // se puede, pero no es necesario
}