package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("csh cli")

	runCmd("echo hi && echo bye")
}

func runCmd(cmd string) {
	lines := strings.Split(cmd, " && ")

	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")

		command := exec.Command(line[0], line[1:]...)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}
