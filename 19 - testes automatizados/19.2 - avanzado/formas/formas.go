package formas

import (
	"math"
)

// la implementacion de la interfaz es implicita en go
type Forma interface {
	Area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Radio float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}
