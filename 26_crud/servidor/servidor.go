package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var insertStatement string = "insert into usuarios (nome, email) values (?, ?)"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	erro = json.Unmarshal(corpoRequisicao, &usuario)

	if erro != nil {
		w.Write([]byte("Falha ao converter o corpo da requisição em JSON"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(insertStatement)
	if err != nil {
		w.Write([]byte("Falha ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Falha ao executar o statement"))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Falha ao obter o ID inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Falha ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Falha ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Falha ao converter os usuários para JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("ID inválido"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linha, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Falha ao buscar o usuário"))
		return
	}
	defer linha.Close()

	var usuario usuario

	if linha.Next() {
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Falha ao escanear o usuário"))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Falha ao converter o usuário para JSON"))
		return
	}

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("ID inválido"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Falha ao converter o corpo da requisição em JSON"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Falha ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.Write([]byte("Falha ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("ID inválido"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Falha ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		w.Write([]byte("Falha ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
