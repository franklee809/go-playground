package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, string := range strings {
		floatPrice, err := strconv.ParseFloat(string, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float.")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
