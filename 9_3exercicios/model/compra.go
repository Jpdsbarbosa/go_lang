package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItemDaCompra
}

type ItemDaCompra struct {
	Nome string
}

func NewCompra(Mercado string, Data time.Time, NomeDosItens []string) (*Compra, error) {

	if Mercado == "" {
		return nil, errors.New("Mercado é obrigatório")
	}

	if len(NomeDosItens) == 0 {
		return nil, errors.New("Itens são obrigatórios")
	}

	var Itens []ItemDaCompra

	for _, nome := range NomeDosItens {
		Itens = append(Itens, ItemDaCompra{Nome: nome})
	}

	return &Compra{Mercado: Mercado, Data: time.Now(), Itens: Itens}, nil

}
