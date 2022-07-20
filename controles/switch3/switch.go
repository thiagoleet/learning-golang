package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "tipo desconhecido"
	}
}

func main() {
	fmt.Println(tipo(3.2))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Olá Mundo"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))

}
