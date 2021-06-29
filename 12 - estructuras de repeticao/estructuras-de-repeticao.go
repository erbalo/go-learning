package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estructuras de repeticao")

	i := 0

	for i < 2 {
		i++
		fmt.Println("Incrementando i")
		//time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 5 { // la variable j solo esta disponible dentro del bloque for (scope)
		fmt.Println("Incrementando j ", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"erick", "dave", "mario"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra, string(letra)) // letra será el código de ascii
	}

	usuario := map[string]string{ // no funciona sobre structs
		"nome":      "erick",
		"sobrenome": "barrera",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	/* FOR infinito
	for true {
		fmt.Println("executando infinito")
	}

	for {
		fmt.Println("executando infinito")
	}
	*/
}
