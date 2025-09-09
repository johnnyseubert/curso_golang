package main

import "fmt"

func main() {

	var nome string = "Johnny"
	sobrenome := "Seubert"

	nomeMae, nomePai := "Magali", "Edilson"

	var idade int = 23
	quantidadeDeGatos := 2

	var altura float64 = 1.74
	peso := 70.5

	var estaTrabalhando bool = true
	estaEstudando := true

	fmt.Println("Nome da mãe:", nomeMae)
	fmt.Println("Nome do pai:", nomePai)

	fmt.Println("Nome:", nome)
	fmt.Println("Sobrenome:", sobrenome)
	fmt.Println("Idade:", idade)
	fmt.Println("Quantidade de gatos:", quantidadeDeGatos)
	fmt.Println("Altura:", altura)
	fmt.Println("Peso:", peso)
	fmt.Println("Está trabalhando?", estaTrabalhando)
	fmt.Println("Está estudando?", estaEstudando)
}
