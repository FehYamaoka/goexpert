package main

func main() {

	// Memória -> Endereço -> Valor

	// variável -> ponteiro que tem um endereço na memória que tem um -> valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(a)
}
