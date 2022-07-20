package main

import (
	"fmt"
	"time"
)

func notaParaConceito(nota float64) string {
	var n = int(nota)
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case nota < 9 && nota >= 7:
		return "B"
	case nota < 7 && nota >= 5:
		return "C"
	case nota < 5 && nota >= 3:
		return "D"
	case nota < 3 && nota >= 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	t := time.Now()

	// switch sem expressão
	// pode ser switch true {}
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
