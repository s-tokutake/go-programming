package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	flag.Parse()
	args := flag.Args()

	for _, arg := range args {
		param := strings.Split(arg, "=")
		go listen(param[0], param[1])
		fmt.Print(arg + " start")
	}

	for {

	}
}

func listen(timezone string, server string) {

	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, timezone)
	}
}

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	jst, _ := time.LoadLocation(timezone)
	for {
		_, err := io.WriteString(c, time.Now().In(jst).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
