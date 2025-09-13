package main

func inverterSinal(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	inverterSinal(&numero)
	println(numero)
}
