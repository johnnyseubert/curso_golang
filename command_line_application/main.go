package main

import (
	"linha_comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal("Erro ao iniciar o aplicativo:", erro)
	}
}
