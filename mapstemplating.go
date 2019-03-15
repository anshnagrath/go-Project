package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	data := map[string]int{"test": 12, "if": 45, "this": 3, "works": 324}
	glob, err := template.ParseGlob("*.html")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(glob)

	err = glob.ExecuteTemplate(os.Stdout, "slice.html", data)
	if err != nil {
		log.Fatalln(err)
	}

}
