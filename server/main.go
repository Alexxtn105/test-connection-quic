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
	fmt.Println("running server...")
	listener, err := quic.ListenAddr("localhost:4242", generateTLSConfig(), nil)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			log.Fatal("Failed to accept session:", err)
		}

		go func() {
			for {
				stream, err := sess.AcceptStream(context.Background())
				if err != nil {
					log.Fatal("Failed to accept stream:", err)
				}

				// эхо полученных данных обратно клиенту
				_, err = io.Copy(stream, stream)
				if err != nil {
					log.Fatal("Failed to echo data:", err)
				}
			}
		}()
	}
}

func generateTLSConfig() *tls.Config {
	//key, cert := generateKeys() // Допустим, что функция generateKeys генерирует TLS ключ и сертификат
	//return &tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	NextProtos:   []string{"quic-echo-example"},
	//}

	//key, cert := generateKeys() // Допустим, что функция generateKeys генерирует TLS ключ и сертификат
	return &tls.Config{
		Certificates: []tls.Certificate{},
		NextProtos:   []string{"quic-echo-example"},
	}
}
