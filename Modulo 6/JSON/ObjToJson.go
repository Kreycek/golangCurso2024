package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Para exportar para JSON o nome das variáveis precisam estar maiusculas
type ss struct {
	//O Hifem indica que não é para mapear o campo
	Nome     string `json:"-"`
	Endereco string `json:"endereco"`
}

func main() {

	c2 := map[string]string{
		"Nome":     "Teste",
		"Endereco": "Teste 2"}

	// dsd := ss{"Ricardo", "teste"}

	pessoa, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	} else {

		// var fdf []uint8 = pessoa
		//retorna bytes
		fmt.Println(bytes.NewBuffer(pessoa))
	}

}
