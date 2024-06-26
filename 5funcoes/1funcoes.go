package main

import "fmt"

func imprimeMensagem(mensagem string) {

	mensagem += ", bom dia!"
	fmt.Println(mensagem)

}

func main() {

	imprimeMensagem("mensagem x")
	imprimeMensagem("mensagem y")

}
