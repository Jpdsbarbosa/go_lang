package main

import "fmt"

func main() {

	x := 5
	y := &x // atribuindo y ao local de memória do x
	*y = 10 // alterando o valor do ponteiro

	imprimirValores(&x, y)
	fmt.Println(x, *y)
	fmt.Println(&x, y)
}

func imprimirValores(x *int, y *int) {

	*x = 20

}
