package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}}`
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	println("")
	println("master:", master)
	println("overlay:", overlay)


	masteredoverlayed := `Names:{{block "list1" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}} {{define "list1"}} {{join . ", "}}{{end}}`
	masteredoverlayedTmpl, err := template.New("masteredoverlayed").Funcs(funcs).Parse(masteredoverlayed)
	if err != nil {
		log.Fatal(err)
	}

	if err := masteredoverlayedTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
6

}

/* 
ralph@ubuntu:~/gowork/src/github.com/GoLangTMcleod/golang-web-dev/002b_templateblock_golang/01$ go run main.go
Names:
- Gamora
- Groot
- Nebula
- Rocket
- Star-Lord
Names: Gamora, Groot, Nebula, Rocket, Star-Lord
 */

/* 
 func (*Template) Clone 
 func (t *Template) Clone() (*Template, error)
 Clone returns a duplicate of the template, including all
 associated templates. The actual representation is not 
 copied, but the name space of associated templates is, 
 so further calls to Parse in the copy will add templates 
 to the copy but not to the original. Clone can be used 
 to prepare common templates and use them with variant 
 definitions for other templates by adding the variants 
 after the clone is made. 
 */

/* 
Nested template definitions 
When parsing a template, another template may be defined and associated with the template being parsed. Template definitions must appear at the top level of the template, much like global variables in a Go program. 
The syntax of such definitions is to surround each template declaration with a "define" and "end" action. 
The define action names the template being created by providing a string constant. Here is a simple example: 

`{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}`

This defines two templates, T1 and T2, and a third T3 that invokes the other two when it is executed. Finally it invokes T3. If executed this template will produce the text 
ONE TWO
 */