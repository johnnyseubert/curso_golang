package main

import "fmt"

func soma(numeros ...int) {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	println(total)

}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	soma(1, 2, 3, 4, 5)
	escrever("Número:", 1, 2, 3, 4, 5)
}
