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

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"McLeod",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/009_predefined-global-functions/01_index/02# go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>predefined global functions</title>
</head>
<body>

zero
one
two
McLeod

<!--
index
a func you can use in a template
it is a "predefined global function"

Index returns the result of indexing its first argument by the
following arguments. Thus "index x 1 2 3" is, in Go syntax,
x[1][2][3]. Each indexed item must be a map, slice, or array.

https://godoc.org/text/template#hdr-Functions
-->

<!-- FYI -->
This is a go template comment





</body>
</html>
 */