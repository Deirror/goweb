package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	addr := flag.String("addr", ":3000", "TCP Server Port")

	go func() {
		time.Sleep(4 * time.Second)
		client := NewFileClient(*addr)

		if err := client.SendFile(4000); err != nil {
			log.Fatal(err)
		}
	}()

	server := NewFileServer(*addr)
	log.Fatal(server.ListenAndStart())

}
