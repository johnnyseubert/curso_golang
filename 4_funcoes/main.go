package main

import "fmt"

func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

func calculosMatematicos(numero1 int8, numero2 int8) (int8, int8, float64, float64) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	multiplicacao := float64(numero1) * float64(numero2)
	divisao := float64(numero1) / float64(numero2)

	return soma, subtracao, multiplicacao, divisao
}

func main() {
	resultado := somar(10, 20)
	fmt.Println("SOMA:", resultado)

	// Funcao anonima
	var function = func(str string) {
		fmt.Println("Função anônima:", str)
	}
	function("Olá, mundo!")

	// Multiplos retornos
	soma, _, multiplicacao, _ := calculosMatematicos(10, 20)

	fmt.Println("SOMA:", soma)
	fmt.Println("MULTIPLICAÇÃO:", multiplicacao)
}
