package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/008_func/02_date-formatting# go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>functions</title>
</head>
<body>

Without date formatting: 2018-10-08 02:57:17.431710937 -0700 PDT m=+0.000739643

With date formatting: 10-08-2018


</body>
</html>
 */