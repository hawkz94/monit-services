package main

import (
	"fmt"
	feature "hello/features"
	"os"
	"time"
)

const delay = 5

func main() {
	showMenu()
	command := readOption()
	callFeature(command)
}

func getUrls() []string {
	return []string{"https://www.alura.com.br/",
		"https://www.alura.com.br/kokwoksksopskopsk",
		"https://www.caelum.com.br/"}
}

func showMenu() {
	println("#############################")
	println("1 - Iniciar monitoramento")
	println("2 - Exibir logs")
	println("3 - Listar diretório")
	println("0 - Fechar aplicação")
	println("#############################")
}

func readOption() int {
	var command int
	fmt.Scan(&command)

	return command
}

func callFeature(option int) {
	switch option {
	case 1:
		urls, err := feature.ReadUrls()

		if err != nil {
			return
		}

		for i := 0; i <= 5; i++ {
			for _, site := range urls {
				feature.Check(site)
			}
			println("")
			time.Sleep(5 * time.Second)
		}

	case 2:
		println("Exibindo logs...")
		feature.ShowLogs()
	case 3:
		feature.ExecList()
	case 0:
		println("Finalizado.")
		os.Exit(0)
	default:
		println("Nenhum comando encontrado.")
	}
}
