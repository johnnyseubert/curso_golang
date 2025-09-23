package main

import "fmt"

func main() {
	canal := escrever("Texto")

	for valor := range canal {
		fmt.Println(valor)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for range 10 {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
		}
		close(canal)
	}()

	return canal
}
