package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	result := map[string]EnvValue{}
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range readDir {
		if file.IsDir() {
			continue
		}
		env, needToRemove := readFileFirstLine(fmt.Sprintf("%s/%s", dir, file.Name()))
		result[file.Name()] = EnvValue{
			Value:      env,
			NeedRemove: needToRemove,
		}
	}

	return result, nil
}

func readFileFirstLine(filePath string) (string, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()
	if firstLine == "" {
		return "", true
	}
	firstLine = strings.Split(firstLine, "\n")[0]
	firstLine = strings.Replace(firstLine, "\u0000", "\n", -1)
	firstLine = strings.TrimRight(firstLine, " \t")

	return firstLine, false
}
