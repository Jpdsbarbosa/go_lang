// funções com retorno unico

package main

import "fmt"

func soma(numero1 int, numero2 int) int {

	resultado := numero1 + numero2
	return resultado

}

func main() {

	resultado := soma(10, 20)
	fmt.Println(resultado)

}
