package main

import "fmt"

type Pessoa struct {
	Nome string
}

type Estudante struct {
	Pessoa
	Matricula string
}

func main() {
	estudante := Estudante{
		Pessoa: Pessoa{
			Nome: "Joana",
		},
		Matricula: "2024001",
	}

	fmt.Println(estudante)

}
