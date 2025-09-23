package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Cria um canal com buffer de 2 posições

	canal <- "Olá mundo!"     // Envia o valor para dentro do canal
	canal <- "Olá novamente!" // Envia outro valor para dentro do canal

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
