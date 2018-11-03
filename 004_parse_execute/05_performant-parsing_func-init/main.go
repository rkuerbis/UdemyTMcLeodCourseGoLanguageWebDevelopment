package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template


//  init() is called before main executes any more program steps
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/004_parse_execute/05_performant-parsing_func-init# go run main.go
******
ONE
************
VESPA
************
TWO
************
ONE
******
 */