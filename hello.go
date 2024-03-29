package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Monitoramentos = 3
	delay          = 5
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	idade := 24
	versao := 1.1

	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() (comando int) {
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{
		"https://www.instagram.com",
		"https://www.alura.com",

		"https://www.facebook.com",
	}
	for i := 0; i < Monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}

		time.Sleep(5 * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registrarLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	arquivo, err := os.ReadFile("sites.txt")
	if err != nil {
		fmt.Println(err)
	}
	sites := strings.Split(string(arquivo), "\r\n")
	return sites
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer arquivo.Close()

	datahora := time.Now().Local().Format("02/01/2006 15:04:05")

	arquivo.WriteString(datahora + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
}

func exibirLogs() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	logs := strings.Split(string(arquivo), "\r\n")
	for _, log := range logs {
		fmt.Println(log)
	}
}
