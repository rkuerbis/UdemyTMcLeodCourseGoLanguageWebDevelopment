package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := region{
		Region: "Southern",
		Hotels: []hotel{
			hotel{
				Name:    "Hotel California",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
			hotel{
				Name:    "H",
				Address: "4",
				City:    "L",
				Zip:     "95612",
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}

/* 
ralph@ubuntu:~/gowork/src/github.com/GoLangTMcleod/golang-web-dev/012_hands-on/04_solution/02$ go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>

Southern

    Hotel California<br>
    42 Sunset Boulevard Los Angeles 95612<br>

    H<br>
    4 L 95612<br>


</body>
</html>
 */