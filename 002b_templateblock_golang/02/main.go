package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	const (
		master  = `Names:{{block "list2" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list2"}} {{join . ", "}}{{end}}`
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

	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}

	
	println("")

	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}

}


/* 
ralph@ubuntu:~/gowork/src/github.com/GoLangTMcleod/golang-web-dev/002b_templateblock_golang/02$ go run main.go
Names: Gamora, Groot, Nebula, Rocket, Star-Lord
Names:
- Gamora
- Groot
- Nebula
- Rocket
- Star-Lord
 */