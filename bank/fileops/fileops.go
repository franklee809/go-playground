package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000.0, errors.New("failed to find new balance file")
	}
	fmt.Println("data :", string(data))
	value, err := strconv.ParseFloat(string(data), 32)

	if err != nil {
		return 1000.0, errors.New("failed to parse value")
	}
	return value, nil
}

func WriteBalanceToFile(balance float64, filename string) {
	var balanceText string = fmt.Sprint(balance)
	fmt.Println(balanceText)
	os.WriteFile(filename, []byte(balanceText), os.FileMode(0644))
}
