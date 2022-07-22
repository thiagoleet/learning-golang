package main

import "fmt"

func obterAprovado(numero int) int {
	defer fmt.Println("fim!")

	const maxValue = 5000

	if numero > maxValue {
		fmt.Println("Valor alto...")
		return maxValue
	}
	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	fmt.Println(obterAprovado(6000))
	fmt.Println(obterAprovado(3000))
}
