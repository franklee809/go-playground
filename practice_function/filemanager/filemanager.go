package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Could not open file!")
		return nil, errors.New("Failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("Failed to read line in file")

	}
	file.Close()
	return lines, nil
}
