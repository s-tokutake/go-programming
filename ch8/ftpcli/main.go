package main

import "net"
import "log"
import "os"
import "bufio"
import "strings"
import "fmt"
import "io"

func main() {
	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		buf := make([]byte, 4096)

		fmt.Println("input = " + text)

		switch {
		case strings.ToUpper(text) == "LS":
			conn.Write([]byte("ls"))
			_, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("err", err)
				return
			}
			fmt.Println(string(buf))

		case strings.ToUpper(text) == "GET":
			conn.Write([]byte("get"))
			_, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("err", err)
				return
			}
			fmt.Println(string(buf))

		default:
			fmt.Println("ERROR: Invalid Command.")
		}

	}
}
