package utils

import (
	"os"
	"os/exec"
)

func ListFiles() ([]byte, error) {
	output, err := exec.Command("ls").Output()
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}

func CreateYaml(fileName string) error {
	fn := fileName + ".yml"
	err := os.WriteFile(fn, []byte{}, 0755)
	if err != nil {
		return err
	}
	return nil
}
