package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"

	// Вариант с методом
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	// Вариант без метода
	fix := strings.ReplaceAll(broken, "#", "o")
	fmt.Println(fix)
}
