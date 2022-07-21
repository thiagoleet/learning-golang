package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	falando := func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}
	go falando()
	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	juntando := func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}
	go juntando()
	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
