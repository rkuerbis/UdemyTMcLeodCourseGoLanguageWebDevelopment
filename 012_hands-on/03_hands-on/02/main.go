package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"
)

type calihotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []calihotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	calihotels := []region {
		region {
			Region: "Northern California",
			Hotels: []calihotel{
        				calihotel{
		        			Name: "Kronborg Inn", 
		        			Address: "1440 Mission Drive", 
		         			City: "Solvang", 
		        			Zip: "93463", 
	         				Region: "Southern", 
		        		},
		        
	        
			        	calihotel{
			         		Name: "Sideways Inn - Buellton", 
			          		Address: "114 East Highway 246", 
			        		City: "Buellton", 
			         		Zip: "93427", 
			        		Region: "Southern", 
			        	},
		         	},
	        	},
		region {
				
			Region: "Central California",
			Hotels: []calihotel{
	         			calihotel{
		         			Name: "Kronborg Inn", 
	         				Address: "1440 Mission Drive", 
		            		City: "Solvang", 
				         	Zip: "93463", 
				        	Region: "Southern", 
		         		},
		
	        			calihotel{
	         				Name: "Sideways Inn - Buellton", 
				            Address: "114 East Highway 246", 
				        	City: "Buellton", 
				        	Zip: "93427", 
					        Region: "Southern", 
				        },
		        	},
				},
	        	
		region {
				
			Region: "Southern California",
			Hotels: []calihotel{
	         			calihotel{
		         			Name: "Kronborg Inn", 
	         				Address: "1440 Mission Drive", 
		            		City: "Solvang", 
				         	Zip: "93463", 
				        	Region: "Southern", 
		         		},
		
	        			calihotel{
	         				Name: "Sideways Inn - Buellton", 
				            Address: "114 East Highway 246", 
				        	City: "Buellton", 
				        	Zip: "93427", 
					        Region: "Southern", 
				        },
		        	},
				},
		}
	

	err := tpl.Execute(os.Stdout, calihotels)
	if err != nil {
		log.Fatalln(err)
	}

	exec.Command("go", " run main.go > main.html")
//	log.Printf("Running command and waiting for it to finish...")

}
