package network

import (
	ws "golang.org/x/net/websocket"
	"log"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Send(ws *ws.Conn, data []byte) {
	b_buf := bytes.NewBuffer([]byte{})
	err := binary.Write(b_buf, binary.BigEndian, int32(len(data)))
	if err != nil {
		log.Fatal(err)
	}
	data = append(b_buf.Bytes(), data...)

	_, err = ws.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %v\n", data)
}

func Recv(ws *ws.Conn) []byte {
	head := make([]byte, 4)
	_, err := ws.Read(head)
	if err != nil {
		log.Fatal(err)
	}
	b_buf := bytes.NewBuffer(head)
	var length int32
	err = binary.Read(b_buf, binary.BigEndian, &length)
	if err != nil {
		log.Fatal(err)
	}

	req_data := make([]byte, length)
	_, err = ws.Read(req_data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %v\n", req_data)
	return req_data
}