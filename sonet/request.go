package sonet

import (
	"github.com/dogchenmao/sogo/soiface"
)

type Request struct {
	connect soiface.IConnect
	data    []byte
}

func (req *Request) GetConnection() soiface.IConnect {
	return req.connect
}

func (req *Request) GetData() []byte {
	return req.data
}
