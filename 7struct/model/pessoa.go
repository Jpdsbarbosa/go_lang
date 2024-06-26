package model

import (
	"time"
)

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

// criando método para alterar dados dentro de uma struct
// neste caso sempre usar ponteiro

func (p *Pessoa) CalculaIdade() {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoDeNascimento

}
