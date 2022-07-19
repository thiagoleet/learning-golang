package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("Linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x) Não funciona

	xs := fmt.Sprint((x))
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)     // sem a necessidade de criar uma nova variável
	fmt.Printf("O valor de x é %f.", x)   // imprimir uma variável do tipo float
	fmt.Printf("O valor de x é %.2f.", x) // imprimir 2 casas decimais (formatado)

	a := 1
	b := 1.999
	c := false
	d := "teste"

	fmt.Printf("%d %f %.1f %t %s", a, b, b, c, d) // cada tipo tem sua %
	fmt.Printf("%v %v %v %v", a, b, c, d)         // %v é genérico

}
