package main

import "fmt"

func inc1(n int) {
	n++
}

// um ponteiro é representado por *
// * é usado para desreferenciar um ponteiro

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n)        // por valor
	fmt.Println(n) // não sofre alterações

	// & é usado para obter o endereço de uma variável
	inc2(&n)       // por referência
	fmt.Println(n) // sofre alterações
}
