package main

import "fmt"

/*

- Exercício 1:
	- Criar um array com 2 posições de inteiros e armazenar uma variavel a soma total da lista.
	- A variável deve ser imprimina no console.
*/

func exercicio1() {

	// fiz dessa forma para caso tivessem mais itens, a soma não estaria presa a 2 itens

	a := 5
	b := 25

	lista := []int{a, b}
	somaLista := 0

	for item := 0; item < len(lista); item++ {
		somaLista += lista[item]
	}

	fmt.Println("Soma da Lista:", somaLista)
}

func exercicio2() {

	/*
	   Dao um slice com os itens "2, 8, 3, 10, 5, 4, 7, 9, 1" que vão de 1 a 10,
	   efetuar a soma em duas variáveis, a primeira números de 1 a 5 e a segunda de 6 a 10.
	   imprimindo os dois resultados
	*/

	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	// soma de 1 a 5 = soma1
	soma1 := 0

	// soma de 6 a 10 = soma2
	soma2 := 0

	for i := 0; i < len(lista); i++ {
		if lista[i] <= 5 {
			soma1 += lista[i]
		} else {
			soma2 += lista[i]
		}
	}

	fmt.Println("Soma de 1 a 5:", soma1)
	fmt.Println("Soma de 6 a 10:", soma2)

}

func main() {

	exercicio1()
	exercicio2()

}
