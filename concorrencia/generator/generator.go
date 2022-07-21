package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan: canal somente-leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			value := r.FindStringSubmatch(string(html))
			if len(value) > 0 {
				c <- value[1]
			}
			c <- "nada"
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.google.com", "https://www.udemy.com", "https://www.github.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)

}
