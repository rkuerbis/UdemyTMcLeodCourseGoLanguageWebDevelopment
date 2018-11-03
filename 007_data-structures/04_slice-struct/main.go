package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-deata-structures/04_slice-struct# go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
<ul>
    
    <li>Buddha - The belief of no beliefs</li>
    
    <li>Gandhi - Be the change</li>
    
    <li>Martin Luther King - Hatred never ceases with hatred but with love alone is healed.</li>
    
    <li>Jesus - Love all</li>
    
    <li>Muhammad - To overcome evil with good is good, to resist evil by evil is evil.</li>
    
</ul>
</body>
</html>
 */