package main

import (
	model "exe82/Model"
	"fmt"
	"time"
)

func main() {

	var itens []model.ItemDaCompra
	itens = append(itens, model.ItemDaCompra{Nome: "Arroz"})
	itens = append(itens, model.ItemDaCompra{Nome: "Feijão"})
	itens = append(itens, model.ItemDaCompra{Nome: "Carne"})

	compra := model.Compra{Mercado: "Mercado X", Data: time.Now(), Itens: itens}
	fmt.Println(compra)

}
