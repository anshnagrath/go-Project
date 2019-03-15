package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := "ansh"
	template := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello world!</title>
	</head>
	<body>
	<p> Hi there my name is ` + name + `</p>
	</body>
	</html>
	`
	fmt.Println(template)
	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(template))
}
