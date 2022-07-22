package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		// condição de parada
		return 1
	default:
		// recursividade
		return n * fatorial(n-1)
	}
}

func main() {
	// sem erro
	resultado := fatorial(5)
	fmt.Println(resultado)

	// com erro
	//resultado2 := fatorial(-4)
	//fmt.Println(resultado2)

}
