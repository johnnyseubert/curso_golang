package main

import "fmt"

func main() {
	// Função anônima auto-invocada
	func(numeros ...int) {
		total := 0
		for _, numero := range numeros {
			total += numero
		}
		println(total)
	}(1, 2, 3, 4, 5)

	// Função anônima atribuída a uma variável
	var funcao = func() {
		fmt.Println("Função anônima atribuída a uma variável")
	}
	funcao()

	// Função anônima com parâmetros e valor de retorno
	retorno := func(txt string) string {
		return fmt.Sprintf("Recebido -> %s", txt)
	}("Teste")

	fmt.Println(retorno)

}
