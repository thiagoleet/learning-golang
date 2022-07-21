package main

import (
	"fmt"
	"time"
)

func main() {
	// não existe while em Go
	i := 1

	// sem bloco de incialização
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// com bloco de inicialização
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("\n%d", i)
	}

	// misturando os tipos de for
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	// laços infinitos
	for {
		fmt.Println("Para sempre")
		time.Sleep(time.Second)
	}

}
