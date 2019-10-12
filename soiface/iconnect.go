package soiface

import "net"

type IConnect interface {
	Start()
	Stop()
	GetConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
}

type HandFunc func(*net.TCPConn, []byte, int) error
