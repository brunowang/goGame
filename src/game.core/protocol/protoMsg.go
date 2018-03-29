package protocol

type IProto interface {
	GetErrorFlag() int8
}

type Proto struct {
}

func (Proto) GetErrorFlag() int8 {
	return 0
}

type ErrorProto struct {
	Code int
}

func (ErrorProto) GetErrorFlag() int8 {
	return 1
}

type NullProto struct {
	Proto
}
