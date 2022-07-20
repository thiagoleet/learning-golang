package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		// erro
		return -1, fmt.Errorf("número inválido: %d", n)

	case n == 0:
		// condição de parada
		return 1, nil
	default:
		// recursividade
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	// sem erro
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	// com erro
	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
