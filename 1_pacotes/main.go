package main

import (
	"github.com/badoux/checkmail"
	"github.com/johnnyseubert/curso_golang/pacotes/auxiliar"
)

func main() {
	auxiliar.Escrever()
	auxiliar.Escrever2()
	err := checkmail.ValidateFormat("johnnygmail.com")

	if err != nil {
		println("Email inválido")
	} else {
		println("Email válido")
	}
}
