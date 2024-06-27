package main

import "time"

type compra struct {
	mercado string
	data    time.Time
	itens   []itenDaCompra
}

type itenDaCompra struct {
	nome string
}
