package main

import (
	"fmt"
	"golangestudo/model"
	"time"
)

func main() {

	fmt.Println("iniciando...")

	endereco := model.Endereco{Rua: "Rua X", Numero: 15, Cidade: "Pelotas"}
	pessoa := model.Pessoa{Nome: "João Pedro Barbosa", Endereco: endereco, DataDeNascimento: time.Date(1996, 8, 31, 0, 0, 0, 0, time.Local)}

	fmt.Println(endereco)
	fmt.Println(pessoa)
	pessoa.CalculaIdade()
	fmt.Println(pessoa.Idade)

	// entendendo herança
	automovelMoto := model.Automovel{Ano: 2022, Placa: "XPTO", Modelo: "CG"}
	moto := model.Moto{Automovel: automovelMoto, Cilindradas: 125}

	fmt.Println(moto)
	fmt.Println(moto.Modelo) // consigo acessar Modelo sem passar por automovel

}
