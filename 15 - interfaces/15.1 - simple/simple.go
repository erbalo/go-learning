package main

import (
	"fmt"
	"math"
)

// la implementacion de la interfaz es implicita en go
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	radio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma e %0.2f\n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
