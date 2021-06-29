package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "erick",
		"sobrenome": "barrera",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "erick",
			"ultimo": "sao",
		},
		"curso": {
			"nome": "engenharia",
			"campus": "campus 1",	
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["ultimo"])
	fmt.Println(usuario2["curso"]["nome"])

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string {
		"nome": "virgo",
	}

	fmt.Println(usuario2)

}