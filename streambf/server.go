package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

type FileServer struct {
	addr string
	ln   net.Listener
}

func NewFileServer(addr string) *FileServer {
	return &FileServer{
		addr: addr,
	}
}

func (s *FileServer) ListenAndStart() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.readLoop(conn)
	}
}

func (s *FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)

		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(buf.Bytes())
		fmt.Printf("received %d bytes over the network\n", n)
	}
}
