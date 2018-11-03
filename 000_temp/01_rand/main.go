package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
//	var n int  !=< 0 
    n := 10
//    var i int 
	for i:=1 ; i < 10 ; i++ {
    	x := rand.Intn(n)
	    fmt.Printf("%T %#v\n\n",x,x)
	}

}
//  outputs an int in random fashion up to seed n=10
//  go run main.go
//  int 9

//  int 4

//  int 2

//  int 8

//  int 7

//  int 8

//  int 6

//  int 1

//  int 7


