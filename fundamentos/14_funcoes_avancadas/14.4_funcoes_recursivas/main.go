package main

import (
	"fmt"
	"time"
)

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	menor := fibonacci(posicao - 1)
	maior := fibonacci(posicao - 2)

	fmt.Println("MENOR:", menor, "MAIOR", maior, "POSIÇÃO:", maior+menor)

	return maior + menor
}

func main() {
	posicao := uint(40)
	now := time.Now()

	for i := range posicao {
		fibonacci(i)
	}

	println("Tempo de execução:", fmt.Sprintf("%.2f", time.Since(now).Seconds()), "segundos")
}
