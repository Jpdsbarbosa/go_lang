package main

import (
	model "exe83/Model"
	"fmt"
	"time"
)

func main() {

	var NomeDosItens []string

	NomeDosItens = append(NomeDosItens, "Arroz")
	NomeDosItens = append(NomeDosItens, "Feijão")
	NomeDosItens = append(NomeDosItens, "Carne")
	NomeDosItens = append(NomeDosItens, "Sabonete")

	compra, err := model.NewCompra(`Mercadinho`, time.Now(), NomeDosItens)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compra)
	}
}
