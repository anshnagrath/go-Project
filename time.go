package main

import (
	"log"
	"math"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("time.html"))
}
func monthDayYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"dbl":      double,
	"sqr":      square,
	"root":     sqrrt,
}

func double(x int) int {
	return x + x
}
func square(x float64) float64 {
	return x * x
}
func sqrrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "time.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
