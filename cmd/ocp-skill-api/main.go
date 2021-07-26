package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is Skill API")

	readConfig := func(path string) error {
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

		return error
	}

	for i := 0; i < 10; i++ {
		if error := readConfig("config"); error != nil {
			fmt.Println(error)
		}
	}
}
