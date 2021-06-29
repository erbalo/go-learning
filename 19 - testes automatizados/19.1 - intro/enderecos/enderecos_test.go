package enderecos_test // para poder generar un nuevo paquete diferente y que se tome dentro de la ejecucion

// nativa do go
// si se pone . antes del import se toma el nombre enderecos como paquete y funciones dentro de el
import (
	. "testes-automatizados/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // ---> corren los test en paralelo
	// go test -v ---> da mas visibilidad de los tests
	// go test --cover ---> coverage del código
	// go test --coverprofile cobertura.txt ---> crear un archivo con los detalles del coverage
	// go tool cover --func=cobertura.txt usa el archivo anterior, para saber el nombre de la funcion que no esta cubierta
	// go tool cover --html=cobertura.txt ---> abre un html con detalles del coverage
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "rua"},
		{"Avenida paulista", "avenida"},
		{"praca das rosas", "Tipo invalido"},
		{"Rodovia coco", "rodovia"},
		{"Estrada qualquer", "estrada"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido e diferente do esperado, Esperava %s e recebeu %s",
				cenario.retornoEsperado,
				tipoDeEnderecoRecebido,
			)
		}
	}
}
