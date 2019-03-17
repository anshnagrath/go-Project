package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("methodsInTemplate.html"))
}
func (p Person) AgeDbl() int {
	return p.Age * 2
}
func (p Person) TakesArgs(x int) int {
	return x * 2
}
func main() {
	an := Person{
		"ansh", 24,
	}

	fmt.Println(tpl)
	err := tpl.Execute(os.Stdout, an)
	if err != nil {
		log.Fatalln(err, "shutting down")
	}
}
