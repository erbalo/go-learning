package main

import (
	"fmt"
	"testes-automatizados/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida paulista")
	fmt.Println(tipoEndereco)
}
