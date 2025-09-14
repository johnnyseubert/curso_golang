package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) // Cria um canal

	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da funcao escrever começar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal // Como se fosse um await, espera receber o valor e depois joga na variavel
	// 	fmt.Println(mensagem)

	// 	if !aberto {
	// 		break
	// 	}
	// }

	for mensagem := range canal { // Outra forma de receber valores do canal, mas só funciona se o canal for fechado
		fmt.Println(mensagem)
	}

	fmt.Println("Programa finalizado")

}

func escrever(texto string, canal chan string) {
	for range 5 {
		canal <- texto // Envia o valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(canal) // Fecha o canal para nao ficar esperando algo que nunca vai chegar
}
