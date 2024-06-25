package main

import (
	"fmt"
)

func aprendendoListas() {

	lista := []int{4, 9, 6, 7}
	fmt.Println("Lista: ", lista)
	fmt.Println("Lista posição1: ", lista[0])
	fmt.Println("Lista posição2: ", lista[1])
	fmt.Println("Lista posição3: ", lista[2])
	fmt.Println("Lista posição4: ", lista[3])
	fmt.Println("Tamanho da Lista: ", len(lista))

}

func listasNaoEstaticas() {

	lista := []int{4, 5, 6, 7}
	fmt.Println("Lista Int: ", lista)
	lista = append(lista, 8)
	fmt.Println("Lista Int depois do apend: ", lista)

	listaDeString := []string{"Java", "Go Lang", "c#"}
	fmt.Println("Lista de Strings: ", listaDeString)
	listaDeString = append(listaDeString, "Ruby")
	fmt.Println("Lista de Strings depois do append: ", listaDeString)
}

func aprendendoMake() {

	lista := make([]int, 1)
	lista[0] = 5
	lista = append(lista, 2)
	lista = append(lista, 3)
	fmt.Printf("%T\n", lista)

	somaTotal := 0

	for i := 0; i < len(lista); i++ {
		somaTotal += lista[i]
	}

	fmt.Println("Média:", somaTotal/len(lista))

}

func sublistas() {

	// quais elementos são menores que 5
	lista := []int{2, 10, 9, 4, 5, 8, 1, 3}

	var listaVazia = make([]int, 0)

	for i := 0; i < len(lista); i++ {
		if lista[i] < 5 {
			listaVazia = append(listaVazia, lista[i])
		} else {
			continue
		}
	}

	fmt.Println("Lista menores que 5:", listaVazia)
}

func index() {

	var listaToda = []int{2, 10, 9, 4, 5, 8, 1, 3}
	seguntaLista := listaToda[:3]
	terceiraLista := listaToda[4:]

	fmt.Println(listaToda)
	fmt.Println(seguntaLista)
	fmt.Println(terceiraLista)

}

func Arrays() {

	// a diferença de um array para os dados trabalhados acima (slices) é que o array tem o numero certo de itens
	lista := [10]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}

	fmt.Println("Array: ", lista)

}

func aprendendoMaps() {

	m := map[string]int{"sp": 10000000000, "cg": 900000}
	fmt.Println(m)

	// criando vazia

	mvazia := make(map[string]int)
	mvazia["rs"] = 2000000
	mvazia["pr"] = 3000000

	fmt.Println(mvazia)

	// verificando se chave existe

	if valor, foiEncontrado := m["rj"]; foiEncontrado {
		fmt.Println(valor)
	} else {
		fmt.Println("Chave não existe")
	}

}

func aprendendoRangeComMap() {

	m := make(map[string]int)
	m["sp"] = 10000000
	m["cg"] = 9000000

	for chave, valor := range m {
		fmt.Println("Cidade:", chave, "----", "Habitantes:", valor)
	}

}

func deletandoItemDoMap() {

	m := make(map[string]int)
	m["sp"] = 10000000
	m["cg"] = 9000000
	delete(m, "cg")

}

func main() {

	aprendendoListas()
	listasNaoEstaticas()
	aprendendoMake()
	sublistas()
	index()
	Arrays()
	aprendendoMaps()
	aprendendoRangeComMap()
	deletandoItemDoMap()

}
