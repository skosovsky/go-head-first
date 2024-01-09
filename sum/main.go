package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers, err := getFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Sum: %0.2f\n", sum)
}

func openFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func closeFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func getFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := openFile(fileName)
	if err != nil {
		return nil, err
	}

	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var number float64
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
