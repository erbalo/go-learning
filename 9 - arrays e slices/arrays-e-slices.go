package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)
	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]string{"p1", "p2", "p3", "p4", "p5"}
	fmt.Println(array2)

	array3 := [...]int{1,2,3} // calcula el tamaño pero NO lo deja dinamico
	fmt.Println(array3)

	// más usado en go, parece un array pero no es un array, internamente es un puntero a un array
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "p alterada"
	fmt.Println(slice2) // el slice ya esta modificado, porque tiene un puntero referenciado

	// arrays internos
	fmt.Println("-----------")
	slice3 := make([]float32, 10, 11) // crea un array interno con capacidad 15, pero regresa las primeras 10 posiciones
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 6) // va a calcular la capacidad del arreglo y se supera, dobla la capicidad del valor interno
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5) // si no se especifica la capacidad, será del mismo tamaño inicial
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 6)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}