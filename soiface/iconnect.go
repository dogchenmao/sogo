package soiface

import "net"

type IConnect interface {
	Start()
	SendNotify()
	SendRequest(MsgType, RequestID, Echo uint32, data []byte)
	SendPulse()
	Close()
}

type HandFunc func(*net.TCPConn, []byte, int) error
