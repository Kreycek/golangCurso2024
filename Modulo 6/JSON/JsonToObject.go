package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Para exportar para JSON o nome das variáveis precisam estar maiusculas
type pessoa struct {
	//O Hifem indica que não é para mapear o campo
	Nome     string `json:"-"`
	Endereco string `json:"endereco"`
	Nif      string `json:"nif"`
}

func main() {

	obj := pessoa{"Ricardo", "Rua Lopes Silva", "32348891"}

	bvbv, error := json.Marshal(obj)

	if error != nil {
		log.Fatal(error)
	} else {

		var p1 pessoa
		pessoaJsonString := `{"Nome":"Lucas","Endereco":"Rua São Miguel do Tapuio","Nif":"314567890" }`

		if erro := json.Unmarshal([]byte(pessoaJsonString), &p1); erro != nil {
			fmt.Println(erro)
		} else {
			fmt.Println("Json transformado em objeto 1 ", p1)
		}

		var p2 pessoa
		json.Unmarshal(bvbv, &p2)
		fmt.Println("Json transformado em objeto 2 ", p2)
	}

}
