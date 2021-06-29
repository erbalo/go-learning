package formas_test

import (
	"math"
	. "teste-avanzado/formas"
	"testing"
)

func TestArea(t *testing.T) { // sub - tests
	/*
			=== RUN   TestArea
			=== RUN   TestArea/Rectangulo
			=== RUN   TestArea/Circulo
			--- PASS: TestArea (0.00s)
		    	--- PASS: TestArea/Rectangulo (0.00s)
		    	--- PASS: TestArea/Circulo (0.00s)
			PASS
			ok      teste-avanzado/formas   0.317s
	*/
	t.Run("Rectangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f e diferente da esperada %f", areaRecebida, areaRecebida)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f e diferente da esperada %f", areaRecebida, areaRecebida)
		}
	})
}
