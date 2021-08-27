package feature

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

func ReadUrls() ([]string, error) {
	var lines []string
	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println(err.Error())
		return lines, err
	}

	read := bufio.NewReader(file)
	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)

		lines = append(lines, line)
		if err == io.EOF {
			break
		}

	}

	return lines, nil
}
