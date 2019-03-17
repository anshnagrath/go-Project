package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name  string
	Admin bool
	Eye   bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("ifand.html"))
}
func main() {
	ansh := Person{
		"ansh", true, true,
	}
	ankit := Person{
		"ankit", false, true,
	}
	Mishra := Person{
		"Mishra", true, false,
	}
	s := []Person{ansh, ankit, Mishra}
	fmt.Println(tpl, s)
	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		log.Fatalln(err, "shutting down")
	}
}
