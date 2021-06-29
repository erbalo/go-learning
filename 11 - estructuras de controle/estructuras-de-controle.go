package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terca"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // entra en esta condicion pero salta a la siguiente condicion, en este ejemplo regresa "Segunda"
	case numero == 2:
		diaDaSemana = "Segunda"
		// En go break no existe para switch
	default:
		diaDaSemana = "Numero invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Estructuras de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { // la variable outroNumero solo esta limitada al scope dentro del flujo if, fuera del if no tendrá valor
		fmt.Println("Numero e maior que zero")
	} else {
		fmt.Println("Numero e menor que zero")
	}

	dia := diaDaSemana(30)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
