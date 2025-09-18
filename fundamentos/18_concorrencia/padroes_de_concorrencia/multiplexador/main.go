package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal1 := escrever("Canal 1")
	canal2 := escrever("Canal 2")

	canalDeSaida := multiplexar(canal1, canal2)

	for i := 0; i < 20; i++ {
		fmt.Println(<-canalDeSaida)
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for range 10 {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
		close(canal)
	}()

	return canal
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalDeSaida <- mensagem
			case mensagem := <-canal2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}
