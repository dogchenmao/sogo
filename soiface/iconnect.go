package soiface

import "net"

type IConnect interface {
	Start()
	SendNotify()
	SendRequest()
	SendPulse()
	Close()
	HandlerMessage()
}

type HandFunc func(*net.TCPConn, []byte, int) error
