package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/005_data/01# go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
</head>
<body>
<h1>The meaning of life: 42</h1>
</body>
</html> 
*/