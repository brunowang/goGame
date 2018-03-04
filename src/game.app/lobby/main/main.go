package main

import (
	ws "golang.org/x/net/websocket"
	"game.core/network"
	"net/http"
	"github.com/golang/protobuf/proto"
	"log"
	"fmt"
	"game.app/lobby/protogen"
)

func main() {
	server := network.NewAcceptor()
	RegisterAllHandler()
	server.Accept()
}

func RegisterAllHandler() {
	http.Handle("/echo", ws.Handler(echoHandler))
}

func echoHandler(ws *ws.Conn) {
	req_data := network.Recv(ws)

	preq := &protogen.Test{}
	if err := proto.Unmarshal(req_data, preq); err != nil {
		log.Fatal("unmarshaling error: ", err)
		return
	}
	fmt.Printf("proto: %v\n", preq)

	preq.Label = proto.String("world")
	preq.Type = proto.Int32(56)
	message, err := proto.Marshal(preq)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	network.Send(ws, message)
}