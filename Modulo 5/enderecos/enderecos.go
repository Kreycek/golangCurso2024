package enderecos

import "strings"

// TipoDeEndereco teste
func TipoDeEndereco(endereco string) string {

	tiposvalidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	primeraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposvalidos {
		if tipo == primeraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeraPalavraDoEndereco
	}

	return "Tipo Inválido"

}
