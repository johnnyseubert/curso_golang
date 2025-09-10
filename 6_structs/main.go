package main

import "fmt"

type Usuario struct {
	Nome  string
	Idade uint8
}

type Endereco struct {
	Logradouro string
	Numero     uint8
}

func main() {
	var user Usuario
	user.Nome = "Johnny"
	user.Idade = 30

	var love Usuario = Usuario{
		Nome:  "Maria",
		Idade: 28,
	}

	fmt.Println(user)
	fmt.Println(love)

}
