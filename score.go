package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Score1 int
	Score2 int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("score.html"))
}
func main() {
	an := Person{
		"ansh", 5, 9,
	}

	//s := []Person{ansh, ankit, Mishra}
	fmt.Println(tpl)
	err := tpl.Execute(os.Stdout, an)
	if err != nil {
		log.Fatalln(err, "shutting down")
	}
}
