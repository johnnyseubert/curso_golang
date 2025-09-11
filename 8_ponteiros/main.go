package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 *int = &variavel1
	fmt.Println(variavel1, *variavel2)
	fmt.Println(&variavel1, variavel2)

	fmt.Println("================================")

	//  Ponteiro é uma referencia de memoria
	var variavel3 int
	var ponteiro *int = &variavel3
	fmt.Println(&variavel3, ponteiro)
	funcaoComPonteiro(&variavel3)
}

func funcaoComPonteiro(valor *int) {
	fmt.Println("Dentro da funcao", valor, *valor)
	*valor = 15
	fmt.Println("Dentro da funcao", valor, *valor)
}
