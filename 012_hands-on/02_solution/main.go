package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear              string
	Fall, Spring, Summer semester
}

var tpl *template.Template

var tpl1 *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl1 = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	years := []year{
		year{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		year{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("go", "build", "main.go")
	log.Printf("Running command and waiting for it to finish...")

	err1 := cmd.Run()
	if err1 != nil {
		log.Printf("Command finished with error: %v", err1)

	}
	f, err3 := os.OpenFile("main.html", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err3)
	}

	err5 := tpl1.Execute(f, years)
	if err5 != nil {
		log.Fatalln(err)
	}


	if err4 := f.Close(); err4 != nil {
		log.Fatal(err)
	}


}

/* 
ralph@ubuntu:~/gowork/src/github.com/GoLangTMcleod/golang-web-dev/012_hands-on/02_solution$ go run main.go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>


    Fall 2020-2021<br>
    
    CSCI-40 - Introduction to Programming in Go - 4<br>
    
    CSCI-130 - Introduction to Web Programming with Go - 4<br>
    
    CSCI-140 - Mobile Apps Using Go - 4<br>
    


    Spring 2020-2021<br>
    
    CSCI-50 - Advanced Go - 5<br>
    
    CSCI-190 - Advanced Web Programming with Go - 5<br>
    
    CSCI-191 - Advanced Mobile Apps With Go - 5<br>
    

    Fall 2021-2022<br>
    
    CSCI-40 - Introduction to Programming in Go - 4<br>
    
    CSCI-130 - Introduction to Web Programming with Go - 4<br>
    
    CSCI-140 - Mobile Apps Using Go - 4<br>
    


    Spring 2021-2022<br>
    
    CSCI-50 - Advanced Go - 5<br>
    
    CSCI-190 - Advanced Web Programming with Go - 5<br>
    
    CSCI-191 - Advanced Mobile Apps With Go - 5<br>
    


</body>
</html>
 */