package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is Skill API")
}

func ReadConfig(paths []string) ([]string, error) {
	if len(paths) < 1 {
		return nil, errors.New("empty config list")
	}

	readFile := func(path string) (string, error) {
		file, error := os.Open(path)

		if error != nil {
			fmt.Println("can't open file")
		}
		defer func() {
			if closeErr := file.Close(); closeErr != nil {
				fmt.Println("can't close file")
				error = closeErr
			}
		}()

		data := new(bytes.Buffer)

		if _, error = data.ReadFrom(file); error != nil {
			return "", error
		}

		return string(data.Bytes()), nil
	}

	output := make([]string, 0)

	for _, path := range paths {
		data, error := readFile(path)

		if error != nil {
			fmt.Sprintf("problems with %v, skip", path)
		} else {
			output = append(output, data)
		}
	}

	return output, nil
}