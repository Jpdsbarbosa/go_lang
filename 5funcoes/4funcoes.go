package main

import (
	"fmt"
)

func operacao(numero1 int, numero2 int) (soma int, sub int, div int, mult int) {

	// não preciso mais criar e atribuir tipo (soma :=) como ela ja foi criada no retorno só preciso atribuir
	soma = numero1 + numero2
	sub = numero1 - numero2
	div = numero1 / numero2
	mult = numero1 * numero2

	return // como coloquei os retornos no nome da função n preciso retornar nada aqui
}

func main() {

	soma, sub, div, mult := operacao(7, 9)
	fmt.Println(soma, sub, div, mult)

}
