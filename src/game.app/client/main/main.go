package main

import (
	ws "golang.org/x/net/websocket"
	"fmt"
	"log"
	"game.app/client/protogen"
	"github.com/golang/protobuf/proto"
	"game.core/network"
)

var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:8080/echo"

func main() {
	ws, err := ws.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	test := &protogen.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &protogen.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	message, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	network.Send(ws, message)
	data := network.Recv(ws)
	test = &protogen.Test{}
	err = proto.Unmarshal(data, test)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Printf("proto: %v", test)

	ws.Close()//关闭连接
}