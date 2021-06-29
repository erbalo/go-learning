package enderecos

import "strings"

// TipoDeEndereco verifica se um endereco tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoMin := strings.ToLower(endereco)
	primeraPalabraDoEndereco := strings.Split(enderecoMin, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeraPalabraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeraPalabraDoEndereco
	}

	return "Tipo invalido"
}
