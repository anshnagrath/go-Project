package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("globalfunctions.html"))
}
func main() {
	s := []string{"test", "index", "asdcasd", "casdca"}
	fmt.Println(tpl, s)
	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		log.Fatalln(err, "shutting down")
	}
}
