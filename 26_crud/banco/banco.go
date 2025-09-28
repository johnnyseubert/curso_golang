package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver de conexao com o banco mysql
)

func Conectar() (db *sql.DB, err error) {
	stringConexao := "root:docker@/docker?charset=utf8&parseTime=True&loc=Local" // usuário:senha@/nome_do_banco para mysql

	db, err = sql.Open("mysql", stringConexao)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
