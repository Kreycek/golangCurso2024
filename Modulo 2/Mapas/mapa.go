package main

import "fmt"

func mainD() {

	f := make(map[string]int)

	f["Valor1"] = 10

	fmt.Println(f["Valor1"])

	usuario2 := map[string]map[string]string{

		"nome": {
			"primeiro": "joao",
			"ultimo":   "pedro",
		},
		"curso": {
			"nome":  "Engenharia",
			"campu": "Campus 1",
		},
	}

	//delete uma chave do map
	delete(usuario2, "nome")

	fmt.Println(usuario2)
}
