package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// enviando um valor à um canal (escrita)
	ch <- 1

	// recebendo dados do canal (leitura)
	<-ch

	ch <- 2

	fmt.Println(<-ch)

}
