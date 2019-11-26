package soiface

type IMessage interface {
	GetData() []byte
	GetBytes() []byte
	GetMsgId() uint32
}
