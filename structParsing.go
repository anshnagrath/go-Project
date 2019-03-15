package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type person struct {
	First string
	Last  string
}

func main() {
	data := person{"ansh", "nagrath"}
	data1 := person{"varun", "nagrath"}
	data2 := person{"rasds", "nagrath"}
	data3 := person{"csdcs", "nagrath"}
	bs := []person{data, data1, data2, data3}
	glob, err := template.ParseGlob("*.html")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(glob)

	err = glob.ExecuteTemplate(os.Stdout, "struct.html", bs)
	if err != nil {
		log.Fatalln(err)
	}

}
