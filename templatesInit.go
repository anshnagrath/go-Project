package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("*.html"))
	fmt.Println(temp)
}
func main() {
	err := temp.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
