package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type FileClient struct {
	addr string
}

func NewFileClient(addr string) *FileClient {
	return &FileClient{
		addr: addr,
	}
}

func (c *FileClient) SendFile(size int) error {
	file := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", c.addr)
	defer conn.Close()

	binary.Write(conn, binary.LittleEndian, size)
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}
	fmt.Printf("bytes written %d\n", n)

	return nil
}
