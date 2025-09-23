package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	menor := fibonacci(posicao - 1)
	maior := fibonacci(posicao - 2)

	return maior + menor
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := range 45 {
		tarefas <- i
	}

	close(tarefas)

	for range 45 {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

}
