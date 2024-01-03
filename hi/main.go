package main

import (
	"github.com/skosovsky/go-head-first/greating"
	"github.com/skosovsky/go-head-first/greating/deutsch"
	"github.com/skosovsky/go-head-first/greating/russian"
)

func main() {
	greating.Hello()
	greating.Hi()
	deutsch.Hallo()
	deutsch.GuttenTag()
	russian.Privet()
	russian.Zdravstvuite()
}
