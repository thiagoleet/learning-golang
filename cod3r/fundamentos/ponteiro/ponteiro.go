package main

import "fmt"

func main() {
	i := 1
	var ponteiro *int = nil // nil é nulo

	ponteiro = &i // obtem o endereço da variável
	*ponteiro++   // * desreferencia o valor associado

	// Go não possui aritmética de ponteiros
	// ponteiro++

	fmt.Println(ponteiro, *ponteiro, i, &i)
}
