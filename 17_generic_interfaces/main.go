package main

import "fmt"

func funcaoAceitaQualquerCoisa(x ...interface{}) {
	fmt.Println(x...)
}

func main() {
	funcaoAceitaQualquerCoisa(123, 'a', "B")
	funcaoAceitaQualquerCoisa("Texto", 456)
	funcaoAceitaQualquerCoisa(true, false)
	funcaoAceitaQualquerCoisa(12.45)
}
