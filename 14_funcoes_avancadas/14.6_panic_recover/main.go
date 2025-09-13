package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Recuperando a execucao")
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso", r)
	}
}

func alunoEstaAprovado(nota1, nota2 float64) bool {
	defer recuperarExecucao()
	media := (nota1 + nota2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média nao pode ser 6")
}

func main() {
	println(alunoEstaAprovado(6, 6))
	println("Pós execucao")
}
