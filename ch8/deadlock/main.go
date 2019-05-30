package main

import "fmt"

func main() {

	message := make(chan string)

	message <- "message" // ここでデッドロック

	fmt.Println(<-message)
}
