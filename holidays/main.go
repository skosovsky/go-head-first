package main

import (
	"fmt"
	"log"

	"github.com/skosovsky/go-head-first/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2019) //nolint:gomnd
	if err != nil {
		log.Println(err)
	}
	err = date.SetMonth(5) //nolint:gomnd
	if err != nil {
		log.Println(err)
	}
	err = date.SetDay(27) //nolint:gomnd
	if err != nil {
		log.Println(err)
	}

	fmt.Println(date)
	fmt.Println(date.Year(), date.Month(), date.Day())

	event := calendar.Event{}
	err = event.SetTitle("Mom's birthday")
	if err != nil {
		log.Println(err)
	}
	err = event.SetYear(2019) //nolint:gomnd
	if err != nil {
		log.Println(err)
	}
	err = event.SetMonth(5) //nolint:gomnd
	if err != nil {
		log.Println(err)
	}
	err = event.SetDay(27) //nolint:gomnd
	if err != nil {
		log.Println(err)
	}

	fmt.Println(event)
	fmt.Println(event.Title(), event.Year(), event.Month(), event.Day())
}
