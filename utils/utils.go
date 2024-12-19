package utils

import (
	"os/exec"
)

func ListFiles() ([]byte, error) {
	output, err := exec.Command("ls").Output()
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}
