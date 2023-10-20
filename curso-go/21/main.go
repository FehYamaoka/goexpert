package main

import (
	"fmt"

	"github.com/fehyamaoka/curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println("Carro: ", carro.Andar())

	fmt.Println("Resultado: ", s)
	fmt.Println("Matemática A: ", matematica.A)
}
