package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args
	//args would be a slice and we are taking an argument from that
	fmt.Println("args", args[1])
	name := args[1]
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
	file, err := os.Create("args.html")
	if err != nil {
		fmt.Println("error cuaght", err)
	}
	defer file.Close()
	io.Copy(file, strings.NewReader(template))

}
