package main

import (
	"fmt"
	"time"
)

// Chanel é a forma de comunicação entre as goroutines

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- base * 4
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	fmt.Println("A")

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
