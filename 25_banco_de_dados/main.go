package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "root:docker@/docker?charset=utf8&parseTime=True&loc=Local" // usuário:senha@/nome_do_banco para mysql

	db, erro := sql.Open("mysql", stringConexao)
	// Abre a conexão com o banco de dados
	// Não verifica se a conexão foi bem sucedida
	// Apenas prepara o banco para uso
	// A conexão de fato só é feita quando é feita a primeira consulta
	// Por isso é importante usar o db.Ping() para verificar se a conexão está ativa
	if erro != nil {
		panic(erro)
	}
	defer db.Close()

	// Verifica se a conexão com o banco de dados está ativa
	if erro = db.Ping(); erro != nil {
		panic(erro)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		panic(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
