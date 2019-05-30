package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)

	}
}

func handleConn(c net.Conn) {

	buf := make([]byte, 4096)
	length, _ := c.Read(buf)
	cmd := string(buf[:length])

	fmt.Println("cmd " + cmd)
	switch {
	case cmd == "ls":
		fmt.Println("ls recv")
		out, _ := exec.Command("ls", "-la").CombinedOutput()
		fmt.Println(string(out))
		_, err := io.WriteString(c, string(out))
		if err != nil && err != io.EOF {
			fmt.Println("err", err)
			return
		}

	case cmd == "get":
		fmt.Println("get recv")
		f, readerr := os.Open("./test.txt")
		if readerr != nil && readerr != io.EOF {
			fmt.Println("err", readerr)
			return
		}
		b, readAllerr := ioutil.ReadAll(f)
		if readAllerr != nil && readAllerr != io.EOF {
			fmt.Println("err", readAllerr)
			return
		}
		_, err := io.WriteString(c, string(b))
		if err != nil && err != io.EOF {
			fmt.Println("err", err)
			return
		}

		f.Close()

	case strings.HasPrefix(cmd, "close"):
		// ??

	}

	defer c.Close()
}
