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

	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/007_data-structures/01_slice/02_variable# go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
<ul>
    
    <li>0 - Gandhi</li>
    
    <li>1 - MLK</li>
    
    <li>2 - Buddha</li>
    
    <li>3 - Jesus</li>
    
    <li>4 - Muhammad</li>
    
</ul>
</body>
</html>
 */