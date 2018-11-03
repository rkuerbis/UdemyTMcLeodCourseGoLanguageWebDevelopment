package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}




/* 
ralph@ubuntu:~/gowork/src/github.com/GoLangTMcleod/golang-web-dev/016_building-a-tcp-server-for-http/01$ go run main.go
GET / HTTP/1.1
***METHOD GET
Host: localhost:8080
Connection: keep-alive
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*\\*;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

GET /favicon.ico HTTP/1.1
***METHOD GET
Host: localhost:8080
Connection: keep-alive
Pragma: no-cache
Cache-Control: no-cache
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36
Accept: image/webp,image/apng,image\*,*\*;q=0.8
Referer: http://localhost:8080/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9

*/

/* 
http://localhost:8080

<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>
*/