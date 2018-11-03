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

type restaurant struct {

    Name string
	Menus []menuitems
} 

var restaurantmenulists []restaurant


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	
	restaurantmenulists = []restaurant {
	                        	restaurant {
									Name: "McDonalds",
									Menus: []menuitems {
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
									},
								},	
								restaurant {
									Name: "Wendys",
									Menus: []menuitems {
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
									},
								},	
							
						}

	err := tpl.Execute(os.Stdout, restaurantmenulists)
	if err != nil {
		log.Fatalln(err)
	}
	
	
}
