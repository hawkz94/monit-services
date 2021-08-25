package listDirs

import (
	"fmt"
	"os/exec"
)

func ExecList() {
	cmd := exec.Command("ls")
	stdout, err := cmd.Output()
	fmt.Print(string(stdout))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
