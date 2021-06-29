package main

import "fmt"

func recuperarExecuaco() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execaco")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecuaco()
	defer fmt.Println("Otro defer")
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA E EXATAMENTE 6!") // antes de mandar a llamar panic, siempre se ejecuta primero defer
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 7))
	fmt.Println("Pos execucao")

	// El error puede continuar si es tratado, un panic no puede mata el proceso completamente
}
