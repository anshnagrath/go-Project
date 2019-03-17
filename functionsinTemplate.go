package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name string
	Moto string
}

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("functionsinTemplate.html"))
}
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}
func main() {
	b := sage{
		"Buddha", "The beliefe of no beliefe",
	}
	c := sage{
		"Gandhi", "I am titanium",
	}
	d := sage{
		"john", "poopooppcspdcas",
	}
	bs := []sage{b, c, d}

	err := tpl.ExecuteTemplate(os.Stdout, "functionsinTemplate.html", bs)
	if err != nil {
		log.Fatalln(err)
	}
}
