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

	//sages := map[string]string{"India":"Gandhi", "America":"MLK", "Meditate":"Buddha", "Love":"Jesus", "Prophet":"Muhammad"}

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/007_data-structures/02_map/02_variable# go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
<ul>
    
    <li>America -  MLK</li>
    
    <li>India -  Gandhi</li>
    
    <li>Love -  Jesus</li>
    
    <li>Meditate -  Buddha</li>
    
    <li>Prophet -  Muhammad</li>
    
</ul>
</body>
</html>
 */