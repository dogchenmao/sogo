package sonet

import (
	"github.com/dogchenmao/sogo/soiface"
)

type client struct {
	Addr       string
	Port       int
	Connect    soiface.IConnect
	msgHandler soiface.IMsgHandle
}

func NewClient(addr string, port int, handler soiface.IMsgHandle) *client {
	return &client{
		Addr:       addr,
		Port:       port,
		msgHandler: handler,
	}
}

func (c *client) Start() {
	c.Connect = NewConnect(c.Addr, c.Port, c.msgHandler)
	go c.Connect.Start()
}

func (c *client) SendRequest(msgType, requestID, echo uint32, data []byte) {
	c.Connect.SendRequest(msgType, requestID, echo, data)
}
