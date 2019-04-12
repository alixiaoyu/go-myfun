package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(st io.Writer, ir io.Reader) {
	_, err := io.Copy(st, ir)
	if err != nil {
		log.Fatal(err)
	}
}
