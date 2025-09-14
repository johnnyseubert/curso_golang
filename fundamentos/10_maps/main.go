package main

import (
	"fmt"
)

type Usuario struct {
	nome string
}

func main() {
	fmt.Println("Maps in Go")

	mapper := map[string]Usuario{
		"nome": {
			nome: "João",
		},
		"sobrenome": {
			nome: "Silva",
		},
	}

	var usuario Usuario = mapper["nome"]

	fmt.Println("Usuário:", usuario)
	fmt.Println("Nome:", usuario.nome)

	mapperAninhado := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
	}
	fmt.Println("ORIGINAL:", mapperAninhado)

	mapperAninhado["endereco"] = map[string]string{
		"rua":    "Rua B",
		"numero": "456",
	}

	fmt.Println("ENDEREÇO ADICIONADO:", mapperAninhado)

	delete(mapperAninhado, "nome")

	fmt.Println("REMOVIDO NOME:", mapperAninhado)

}
