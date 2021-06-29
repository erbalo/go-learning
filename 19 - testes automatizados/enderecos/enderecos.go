package enderecos

import "strings"

// TipoDeEndereco verifica se um endereco tem im tipo valid e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estada", "rodovia"}
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
