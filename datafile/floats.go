package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значения float64 из каждой строки файла
func GetFloats(filePath string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var number float64
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
