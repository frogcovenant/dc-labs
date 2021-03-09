package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func read(connection net.Conn, channel chan string) {
	localTime := make([]byte, 64)
	for {
		bytes, errB := connection.Read(localTime)
		if errB != nil {
			log.Fatal(errB)
		}
		channel <- string(localTime[:bytes])
	}
}

func main() {
	channel := make(chan string, len(os.Args[1:]))
	for _, argument := range os.Args[1:] {
		port := strings.Split(argument, ":")[1]
		connection, errC := net.Dial("tcp", "localhost:"+port)
		if errC != nil {
			log.Fatal(errC)
		}
		defer connection.Close()
		go read(connection, channel)
	}
	for msg := range channel {
		fmt.Print(msg)
	}
}
