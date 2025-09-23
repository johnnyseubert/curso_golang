package main

import "fmt"

func calculosMatematicoImplicito(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	divisao := n1 / n2 //Divisao nao é retornado somente utilizado dentro da função
	fmt.Println("Divisão:", divisao)
	// valores nomeados podem ser retornados diretamente com a palavra return
	return
}

func calculosMatematicoRetorno(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2

	return soma, subtracao
}

func main() {
	soma, subtracao := calculosMatematicoRetorno(10, 5)
	println("Soma:", soma)
	println("Subtração:", subtracao)

	somaImplicito, subtracaoImplicito := calculosMatematicoImplicito(10, 5)
	println("Soma Implicito:", somaImplicito)
	println("Subtração Implicito:", subtracaoImplicito)

}
