package main

import (
	"fmt"

	listDirs "hello/features"
)

func main() {

	showMenu()
	command := readOption()
	callFeature(command)
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
		println("2 - Exibir logs")
	case 2:
		println("2 - Exibir logs")
	case 3:
		listDirs.ExecList()
	case 0:
		println("0 - Fechar aplicação")
	default:
		println("Nenhum comando encontrado.")
	}
}
