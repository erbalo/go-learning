package main

import "fmt"

func main() {
	fmt.Println("Funcoes anonimas")

	retorno := func(text string) string {
		return fmt.Sprintf("recebido -> %s %s", "Ola mundo", text)
	}

	fmt.Println(retorno("erick"))

}
