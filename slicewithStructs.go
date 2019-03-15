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
	data4 := person{"ancasdsh", "nagrath"}
	data5 := person{"varucasdn", "nagrath"}
	data6 := person{"rasds", "nagratcasdh"}
	data7 := person{"csdcs", "nagcasdrath"}
	bs2 := []person{data4, data5, data6, data7}
	glob, err := template.ParseGlob("*.html")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(glob)
	pass := struct {
		Bs  []person
		Bs2 []person
	}{
		bs,
		bs2,
	}
	err = glob.ExecuteTemplate(os.Stdout, "slicewithStructs.html", pass)
	if err != nil {
		log.Fatalln(err)
	}

}
