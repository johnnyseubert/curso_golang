package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20

	soma := num1 + num2
	subtracao := num1 - num2
	multiplicacao := num1 * num2
	divisao := num1 / num2
	modulo := num1 % num2

	fmt.Println("SOMA:", soma)
	fmt.Println("SUBTRAÇÃO:", subtracao)
	fmt.Println("MULTIPLICAÇÃO:", multiplicacao)
	fmt.Println("DIVISÃO:", divisao)
	fmt.Println("RESTO DA DIVISÃO:", modulo)

	var numero1 int16 = 1000
	var numero2 int32 = 2000

	soma2 := int32(numero1) + numero2

	fmt.Println("SOMA:", soma2)
}
