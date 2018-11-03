package main

import (
	"log"
	"os"
	"text/template"
)

type items string

type menuitems struct {
	Mealtime string
	Mealitems []items
}

type menus []menuitems

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	
	allmenus := menus {
		       
		    		menuitems {
						Mealtime: "Breakfast",
						Mealitems: []items {
						            	"Two Eggs and Bacon",
										"two Eggs and Ham", 
										"Two Eggs and Bacon and Toast", 
										"Two Eggs and Ham and Toast", 
									 
						            },
						},
					menuitems {	
				    	Mealtime: "Lunch",
				     	Mealitems: []items {
										"Regular Hamburger and Fries",
										"Double Hamburger and Fries", 
										"Triple Hamburger and Fries", 
										"Bacon Lettuce and Tomato Sandwitch", 
									 
						            },
						},
					menuitems {
					    Mealtime: "Dinner",
					    Mealitems: []items {
										"Regular Hamburger and Fries",
										"Double Hamburger and Fries", 
										"Triple Hamburger and Fries", 
										"Bacon Lettuce and Tomato Sandwitch", 
										 
									},
						},
             	}	
	/*					
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
	
    */
	err := tpl.Execute(os.Stdout, allmenus)
	if err != nil {
		log.Fatalln(err)
	}
	
	
}
