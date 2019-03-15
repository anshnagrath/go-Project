package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	data := []string{"test", "if", "this", "works"}
	glob, err := template.ParseGlob("*.html")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(glob)

	err = glob.ExecuteTemplate(os.Stdout, "slice.html", data)
	if err != nil {
		log.Fatalln(err)
	}
	// err = tpl2.ExecuteTemplate(os.Stdout, "index.html", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// nf, err := os.Create("index3.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer nf.Close()
	// 	err = tpl.Execute(nf, nil)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
}
