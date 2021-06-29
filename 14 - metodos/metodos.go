package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// esta funcion esta agrupada en una estructura, todos los usuarios tienen metodo salvar
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados de usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// para hacer alteraciones al struct al ser llamado y no usar copia
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"erick", 27}
	usuario2 := usuario{"dave", 17}
	usuario1.salvar()
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
