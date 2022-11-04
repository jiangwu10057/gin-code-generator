package util

import (
	"fmt"
	"os/exec"
)

func GoFileFormat(file string) error {
	cmd := exec.Command("gofmt", "-l", "-w", file)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
