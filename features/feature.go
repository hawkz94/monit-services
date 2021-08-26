package feature

import (
	"fmt"
	"net/http"
	"os/exec"
)

func Check(url string) {
	fmt.Println("Monitorando...")
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if res.StatusCode == 200 {
		fmt.Println("Site:", url, "carregado com sucesso!")
	} else {
		fmt.Println("Site:", url, "com erro!", res.StatusCode)
	}
}

func ExecList() {
	cmd := exec.Command("ls")
	stdout, err := cmd.Output()
	fmt.Print(string(stdout))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
