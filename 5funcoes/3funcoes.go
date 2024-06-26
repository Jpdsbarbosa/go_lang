// funções que retornam mais de um valor

package main

import (
	"fmt"
)

func operacao(numero1 int, numero2 int) (int, int, int, int) {

	soma := numero1 + numero2
	sub := numero1 - numero2
	div := numero1 / numero2
	mult := numero1 * numero2

	return soma, sub, div, mult
}

func main() {

	soma, sub, div, mult := operacao(1, 2)
	fmt.Println(soma, sub, div, mult)

}
