package main

import (
	"fmt"
)

func main() {

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad"}

	for k, v := range sages {
		fmt.Println(k, " - ", v)
	}

}

/* 
/home/ralph/gowork/src/github.com/GoLangTMcleod/golang-web-dev/007_data-structures/02_map/02_variable/no-order# go run main.go
India  -  Gandhi
America  -  MLK
Meditate  -  Buddha
Love  -  Jesus
Prophet  -  Muhammad

 */
