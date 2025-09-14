package main

import "fmt"

var numero int

func init() {
	fmt.Println("Executando a função init")
	fmt.Println("Número:", numero)
	numero = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println("Número:", numero)
}
