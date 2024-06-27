package main

import (
	"fmt"
	"time"
)

func main() {

	var itens []itenDaCompra
	itens = append(itens, itenDaCompra{nome: "Arroz"})
	itens = append(itens, itenDaCompra{nome: "Feijão"})
	itens = append(itens, itenDaCompra{nome: "Carne"})

	compra := compra{mercado: "Mercado X", data: time.Now(), itens: itens}
	fmt.Println(compra)

}
