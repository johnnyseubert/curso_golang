package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"-"` // Simbolo - ignora o campo na hora de converter para json
}

func main() {
	// Marshal é o processo de converter para json Unmarshal é o processo de converter de json para struct
	c := Cachorro{"Rex", "Dalmata", 3}

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
		return
	}
	fmt.Println(string(cachorroEmJSON))

	c2 := map[string]string{
		"nome":  "Tob",
		"raca":  "Pastor Alemao",
		"idade": strconv.Itoa(5),
	}
	cachorroEmJSON2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
		return
	}
	fmt.Println(string(cachorroEmJSON2))

	// ======================= UNMARSHAL =======================

	cachorroUnmarshal := `{"nome":"Bob","raca":"Labrador","idade":4}`

	var c3 Cachorro

	if erro := json.Unmarshal([]byte(cachorroUnmarshal), &c3); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c3)

}
