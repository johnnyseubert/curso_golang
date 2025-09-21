package main

import (
	"fmt"
	"testes/enderecos"
)

// go test --cover
// go test ./...
// go tool cover --html=cobertura.txt

/*
	Caso queira gerar um arquivo de cobertura de testes e mostrando o que nao esta coberto
	go test --coverprofile cobertura.txt
	go tool cover --html=cobertura.txt
*/

func main() {
	result := enderecos.TipoDeEndereco("Avenida Paulista")

	fmt.Print(result)
}
