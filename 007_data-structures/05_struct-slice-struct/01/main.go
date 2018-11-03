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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/007_data-structures/05_struct-slice-struct/01# go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Wisdom & Transport</title>
</head>
<body>

<ul>
    
    <li>Buddha - The belief of no beliefs</li>
    
    <li>Gandhi - Be the change</li>
    
    <li>Martin Luther King - Hatred never ceases with hatred but with love alone is healed.</li>
    
</ul>

<ul>
    
    <li>Ford - F150 - 2</li>
    
    <li>Toyota - Corolla - 4</li>
    
</ul>

</body>
</html>
 */