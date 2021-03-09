// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	for {
		location, errZ := time.LoadLocation(timeZone)
		if errZ != nil {
			return // e.g., timezone not parsed correctly :(
		}
		_, errW := io.WriteString(c, timeZone+"\t : "+time.Now().In(location).Format("15:04:05\n"))
		if errW != nil {
			return // e.g., client disconnected :(
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	timeZone := os.Getenv("TZ")
	port := flag.String("port", "4200", "This is used for the port number")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, timeZone) // handle connections concurrently
	}
}
