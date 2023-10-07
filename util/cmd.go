package util

import (
	"os/exec"
)

func Command(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	// 此处是windows版本
	// c := exec.Command("cmd", "/C", cmd)
	_, err := c.CombinedOutput()
	return err
}
