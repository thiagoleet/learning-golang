package main

import "fmt"

func main() {
	funcionariosESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcionariosESalarios["Rafael Filho"] = 1350.0
	delete(funcionariosESalarios, "Inexistente")

	for nome, salario := range funcionariosESalarios {
		fmt.Println(nome, salario)
	}
}
