package main

import (
	"log"
	"os"
	"os/exec"	
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p := person{
		Name: "Ian Fleming",
		Age:  56,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("go", " run main.go > main.html")
	log.Printf("Running command and waiting for it to finish...")

	err1 := cmd.Run()
	log.Printf("Command finished with error: %v", err1)


}

// Generally speaking, best practice:
// call functions in templates for formatting only; not processing or data access.

// The main reasons you don't want to do any data processing in your template:
// (1) separation of concerns
// (2) if you're using a function more than once in a template,
// the server needs to do the processing more than once.
// (though the standard library might cache processing -
// I've yet to dig into this - if you find the answer, let me know)

/* 
ralph@ubuntu:~/gowork/src/github.com/GoLangTMcleod/golang-web-dev/011_composition-and-methods/04_method$ go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Composition</title>
</head>
<body>

<h1>Ian Fleming</h1>
<h2>56</h2>
<h3>7</h3>
<h3>112</h3>
<h3>112</h3>
<h3>224</h3>
</body>
</html>
 */