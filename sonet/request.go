package sonet

import (
	"github.com/dogchenmao/sogo/soiface"
)

type Request struct {
	connect soiface.IConnect
	msg     soiface.IMessage
}

func (req *Request) GetConn() soiface.IConnect {
	return req.connect
}

func (req *Request) GetMsg() soiface.IMessage {
	return req.msg
}
