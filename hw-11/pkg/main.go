package pkg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(args []string) {
	path := args[0]

	envs, err := ReadDir(path)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = RunCmd(args, envs)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func ReadDir(dir string) (map[string]string, error) {
	envMap := make(map[string]string)

	files, err := filepath.Glob(dir)
	if err != nil {
		return envMap, err
	}

	for _, file := range files {

		file_from, err := os.OpenFile(file, os.O_RDWR, 0644)
		if err != nil {
			if os.IsNotExist(err) {
				return envMap, errors.New("no such file")
			}
		}
		defer file_from.Close()

		data := make([]byte, 64)
		offset := 0

		for {
			read, err := file_from.Read(data)
			offset += read
			if err == io.EOF {
				break
			}
		}
		envMap[filepath.Base(file)] = string(data[:offset])
	}

	return envMap, nil
}

func RunCmd(args []string, env map[string]string) error {
	command := args[1]
	currentArgs := args[2:]

	for key, value := range env {
		os.Setenv(key, value)
	}

	c1 := exec.Command(command, currentArgs...)

	c1.Stdout = os.Stdout

	if err := c1.Start(); err != nil {
		return err
	}
	if err := c1.Wait(); err != nil {
		return err
	}

	return nil
}
