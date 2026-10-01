package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
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
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.Contains(name, "=") {
			continue
		}

		envValue, e := readFromFile(filepath.Join(dir, name))
		if e != nil {
			return nil, e
		}

		env[name] = envValue
	}

	return env, nil
}

func readFromFile(filename string) (EnvValue, error) {
	file, err := os.Open(filename)
	if err != nil {
		return EnvValue{}, err
	}

	defer file.Close()

	var firstLine string
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine = scanner.Text()
	}

	if err = scanner.Err(); err != nil {
		return EnvValue{}, err
	}

	fixedTerminalLine := bytes.ReplaceAll([]byte(firstLine), []byte{0x00}, []byte{'\n'})

	finalLine := strings.TrimRight(string(fixedTerminalLine), " \t")

	if finalLine == "" {
		return EnvValue{NeedRemove: true}, nil
	}

	return EnvValue{Value: finalLine}, nil
}
