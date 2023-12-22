package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Douglas"
	idade := 24
	versao := 1.1

	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("o tipo da variavel é", reflect.TypeOf(nome))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(comando)

	fmt.Println("O comando escolhido foi", comando)

}
