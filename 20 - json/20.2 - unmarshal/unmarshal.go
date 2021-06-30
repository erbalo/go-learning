package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"-"` // si se pone guion (tag) ignora el campo
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroJson := `{"nome":"lorenza","raca":"mestizo","idade":3}`

	var c cachorro
	if erro := json.Unmarshal([]byte(cachorroJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachoroJson2 := `{"nome":"tobby","raca":"poddle"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachoroJson2), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
