# Curso golang

### Modulos
- No `go` para poder trabalhar com múltiplos arquivos e pacotes é necessário que existe um módulo raiz entao no inicio do projeto rodar o comando `go mod init github.com/johnnyseubert/pacote`


---

### Exportar e importar
- Como o `go` não é uma linguagem orientada a objetos, não existe os modificadores de acesso `public` `private` `protected` etc...
- Para isso as `variaveis, structs, interfaces, funções etc..` dependem do nome, se a primeira letra for `Maiúscula` o item é public se for `Minúscula` é privado

<br/>
<br/>

```go
package auxiliar
func escreverPrivado(){}
func EscreverPublico(){}


package main
func main(){
   auxiliar.escreverPrivado() // <- vai gerar um erro pois está privado
   auxiliar.EscreverPublico() // <-vai funcionar pois está publico
}
```
