package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup //Grupo de espera

	waitGroup.Add(2) // Diz quantas goroutines vão ser esperadas

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() // Diz que uma das goroutines terminou
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // Diz que uma das goroutines terminou
	}()

	// Garantir que o sistema espere as goroutines finalizarem
	waitGroup.Wait()
}

func escrever(texto string) {
	for range 5 {
		fmt.Println(texto)
		time.Sleep(time.Millisecond * 250)
	}
}
