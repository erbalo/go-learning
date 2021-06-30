package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"lorenza", "mestizo", 3}
	fmt.Println(c)

	jsonStr, erro := json.Marshal(c) // striuct, map -> json

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(jsonStr)
	fmt.Println(bytes.NewBuffer(jsonStr))

	c2 := map[string]string{
		"nome": "tobby",
		"raca": "poddle",
	}

	jsonStr2, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(jsonStr2))
}
