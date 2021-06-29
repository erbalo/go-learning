package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resutaldo sera retornado", n1, n2)
	fmt.Println("Entrando a funcao pra verificar se o aluno ta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Funcoes defer")
	//defer funcao1()
	// defer -> adiar
	//funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
