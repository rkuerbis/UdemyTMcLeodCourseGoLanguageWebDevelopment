package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}

/* 
ralph@ubuntu:~/gowork/src/github.com/GoLangTMcleod/golang-web-dev/015_understanding-TCP-servers/01_write$ go run main.go
^Csignal: interrupt
ralph@ubuntu:~/gowork/src/github.com/GoLangTMcleod/golang-web-dev/015_understanding-TCP-servers/01_write$ 
 */

/* 
ralph@ubuntu:~$ telnet localhost 8080
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.

Hello from TCP server
How is your day?
Well, I hope!Connection closed by foreign host.
ralph@ubuntu:~$ 
 */