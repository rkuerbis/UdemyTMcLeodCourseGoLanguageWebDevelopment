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

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/009_predefined-global-functions/03_comparison# go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <m/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/009_predefined-global-functions/03_comparison# go run main.go
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>predefined global functions</title>
	</head>
	<body>
	
	
	
	
	SCORE 1 < SCORE 2
	
	
	
	
	
	LAST ONE: SCORE 1 < 10
	
	</body>
	</html>eta charset="UTF-8">
    <title>predefined global functions</title>
</head>
<body>




SCORE 1 < SCORE 2





LAST ONE: SCORE 1 < 10

</body>
</html>
 */