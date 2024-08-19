package main

import "fmt"

// Garante que você só recebe nesse canal
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// Garante que você apenas lê nesse canal
func ler(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
