package main

import "fmt"

func evenOrOdd(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d is %s\n", i, evenOrOdd(i))
	}
}
