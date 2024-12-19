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

func CreateYaml(fileName string) ([]byte, error) {
	fn := fileName + ".yml"
	cmd := exec.Command("touch", fn)
	output, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}
