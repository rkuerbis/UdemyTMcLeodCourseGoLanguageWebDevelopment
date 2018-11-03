package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"
)
/* 
func (c *closeOnce) finalclose() {
	c.exec.close()
}
 */
type calihotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template
var tpl1 *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl1 = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	calihotels := []calihotel{
		calihotel{
			Name:    "Kronborg Inn",
			Address: "1440 Mission Drive",
			City:    "Solvang",
			Zip:     "93463",
			Region:  "Southern",
		},

		calihotel{
			Name:    "Sideways Inn - Buellton",
			Address: "114 East Highway 246",
			City:    "Buellton",
			Zip:     "93427",
			Region:  "Southern",
		},
	}

	err := tpl.Execute(os.Stdout, calihotels)
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

	err5 := tpl1.Execute(f, calihotels)
	if err5 != nil {
		log.Fatalln(err)
	}


	if err4 := f.Close(); err4 != nil {
		log.Fatal(err)
	}


	/* 
	cmd2 := exec.Command("./main")
	log.Printf("Running command and waiting for it to finish...")

	err2 := cmd2.Run()
	log.Printf("Command finished with error: %v", err2)
 */
	/* 
	err2 := exec.Close()
	log.Printf("Command finished with close error: %v", err2)
    , ">", "/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/012_hands-on/03_hands-on/01/main.html"
	*/



}
