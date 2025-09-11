package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(nota1, nota2 float64) bool {
	defer fmt.Println("Média calculada, resultado sera retornado")

	fmt.Println("Calculando a média do aluno")
	media := (nota1 + nota2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1()
	funcao2()

	println(alunoEstaAprovado(7, 8))

}
