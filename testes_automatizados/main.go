package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {
	result := enderecos.TipoDeEndereco("Avenida Paulista")

	fmt.Print(result)
}
