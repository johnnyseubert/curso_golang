package main

import "fmt"

func main() {
	numero := 1

	dia := diaDaSemana(numero)

	fmt.Printf("Dia da semana é [%s]", dia)

}

func diaDaSemana(dia int) string {
	var diaSemana string

	switch dia {
	case 1:
		diaSemana = "Domingo"
		fallthrough // Continua para o proximo case ignorando a verificação entao mesmo passando o dia 1 ele vai entrar no case 1 vai fazer oq precisa e vai entrar no case 2
	case 2:
		diaSemana = "Segunda"
	case 3:
		diaSemana = "Terça"
	case 4:
		diaSemana = "Quarta"
	case 5:
		diaSemana = "Quinta"
	case 6:
		diaSemana = "Sexta"
	case 7:
		diaSemana = "Sábado"
	default:
		diaSemana = "Dia inválido"
	}

	return diaSemana

}
