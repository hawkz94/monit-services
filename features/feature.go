package feature

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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
		registerLog(url, true)
	} else {
		fmt.Println("Site:", url, "com erro!", res.StatusCode)
		registerLog(url, false)
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

func registerLog(url string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + url + " - online: " + strconv.FormatBool(status) + "\n")
}

func ShowLogs() {
	file, err := os.Open("logs.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	read := bufio.NewReader(file)
	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)

		fmt.Println(line)
		if err == io.EOF {
			break
		}
	}
}
