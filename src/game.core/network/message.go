package network

import "game.core/protocol"

type Message struct {
	CmdId       int32
	PushId      int32
	RequestUUID int32
	ErrorFlag   int8
	Body        protocol.IProto
}

func NewMessage(cmdId int32) Message {
	return Message{CmdId:cmdId}
}