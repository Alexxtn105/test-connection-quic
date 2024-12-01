package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go"
	"io"
	"log"
)

func main() {
	session, err := quic.DialAddr(context.Background(), "localhost:4242", &tls.Config{InsecureSkipVerify: true}, nil)
	if err != nil {
		log.Fatal("Failed to dial:", err)
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		log.Fatal("Failed to open stream:", err)
	}

	if _, err := fmt.Fprintf(stream, "Hello, QUIC Server!\n"); err != nil {
		fmt.Println("error formatting stream")
	}
	buf := make([]byte, 1024)
	n, err := io.ReadFull(stream, buf)
	if err != nil {
		log.Fatal("Failed to read from stream:", err)
	}

	fmt.Printf("Server says: %s", string(buf[:n]))
}
