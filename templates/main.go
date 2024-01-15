package main

import (
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func main() {
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.5)

	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", false)

	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do", "re", "mi"})

	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 27})
	executeTemplate(templateText, []float64{})
	executeTemplate(templateText, nil)

	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})

	templateText = "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)
	subscriber = Subscriber{Name: "Joy Carr", Rate: 5.99, Active: false}
	executeTemplate(templateText, subscriber)
}

func executeTemplate(text string, data any) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
