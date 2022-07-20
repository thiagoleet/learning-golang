package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // Continua descendo para os próximos cases
	case 9:
		return "A"
	case 8, 7: // range de opções
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(1.1))
	fmt.Println(notaParaConceito(11.1))
	fmt.Println(notaParaConceito(-5.1))
}
