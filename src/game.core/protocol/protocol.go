package protocol

import "reflect"

type ProtoType int8
const (
	REQ_RESP = iota
	PUSH
	NOTIFY
)

type Protocol struct {
	Id              int32
	Type            ProtoType
	Name            string
	UpstreamType    reflect.Type
	DownstreamType  reflect.Type
	ErrorStreamType reflect.Type
}

func CreateReqRes(id int32, name string, u, d reflect.Type) Protocol {
	return create(id, REQ_RESP, name, u, d, reflect.TypeOf("ErrorProto"))
}

func CreatePush(id int32, name string, d reflect.Type) Protocol {
	return create(id, PUSH, name, reflect.TypeOf("NullProto"), d, reflect.TypeOf("NullProto"))
}

func CreateNotify(id int32, name string, u, d reflect.Type) Protocol {
	return create(id, NOTIFY, name, u, reflect.TypeOf("NullProto"), reflect.TypeOf("NullProto"))
}

func create(id int32, typ ProtoType, name string, u, d, e reflect.Type) Protocol {
	return Protocol{
		Id:id,
		Type:typ,
		Name:name,
		UpstreamType:u,
		DownstreamType:d,
		ErrorStreamType:e,
	}
}