package soclient

import (
	"net"

	"github.com/dogchenmao/sogo/soiface"
	"golang.org/x/net/websocket"
)

type Connection struct {
	Conn        *websocket.Conn
	SessionID   uint32
	HandleAPI   soiface.HandFunc
	IsClosed    bool
	Router      soiface.IRouter
	ExitBufChan chan bool
}

func NewConnect(conn *websocket.Conn, sessionID uint32, handle soiface.HandFunc, router soiface.IRouter) *Connection {
	s := &Connection{
		Conn:        conn,
		SessionID:   sessionID,
		HandleAPI:   handle,
		IsClosed:    true,
		Router:      router,
		ExitBufChan: make(chan bool, 1),
	}
	return s
}

func (con *Connection) StartReader() {

}

func (con *Connection) Start() {
	go con.StartReader()
	for {
		select {
		case <-con.ExitBufChan:
			return
		}
	}
}

func (con *Connection) Stop() {
	if con.IsClosed == true {
		return
	}

	con.IsClosed = true

	con.Conn.Close()

	con.ExitBufChan <- true

	close(con.ExitBufChan)
}

func (con *Connection) RemoteAddr() net.Addr {
	return con.Conn.RemoteAddr()
}

func (con *Connection) GetConnID() uint32 {
	return con.SessionID
}

func (con *Connection) GetConnection() *websocket.Conn {
	return con.Conn
}
