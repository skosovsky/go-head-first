package main

import (
	"fmt"
	"github.com/skosovsky/go-head-first/magazine"
)

func main() {
	var s magazine.Subsriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}
