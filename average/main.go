// Package main делает расчет среднего числа из чисел файла
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers, err := getFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}

// getFloats читает значения float64 из каждой строки файла
func getFloats(filePath string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(filePath)
	if err != nil {
		return numbers, err
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil
}
