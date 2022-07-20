package main

import "fmt"

func main() {

	numeros := [...]int{1, 2, 3, 4, 5}

	// com índice
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// sem índice
	for _, num := range numeros {
		fmt.Println(num)
	}

}
