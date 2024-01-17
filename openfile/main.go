package main

import (
	"log"
	"os"
)

func main() {
	option := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("aardvark.txt", option, os.FileMode(0644))
	check(err)
	_, err = file.Write([]byte("amazing!\n"))
	check(err)
	err = file.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
