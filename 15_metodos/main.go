package main

type Usuario struct {
	Nome  string
	Idade uint8
}

func (user Usuario) SalvarDados() {
	println("Salvo no banco de dados o usuario:", user.Nome)
}

// Se for fazer uma alteração na struct, usar ponteiro para persistir o dado fora da funcao
func (user *Usuario) AumentarIdade() {
	user.Idade++
}

func main() {
	usuario := Usuario{"Gustavo", 29}
	println(usuario.Nome, usuario.Idade)

	usuario.SalvarDados()
	usuario.AumentarIdade()
	println(usuario.Nome, usuario.Idade)

}
