package sonet

import (
	"fmt"
	"net"

	"github.com/dogchenmao/sogo/soiface"
)

type Connect struct {
	Conn        *net.TCPConn
	SessionID   uint32
	HandleAPI   soiface.HandFunc
	IsClosed    bool
	Router      soiface.IRouter
	ExitBufChan chan bool
}

func NewConnect(conn *net.TCPConn, sessionID uint32, handle soiface.HandFunc, router soiface.IRouter) *Connect {
	s := &Connect{
		Conn:        conn,
		SessionID:   sessionID,
		HandleAPI:   handle,
		IsClosed:    true,
		Router:      router,
		ExitBufChan: make(chan bool, 1),
	}
	return s
}

func (con *Connect) StartReader() {
	fmt.Println(con.RemoteAddr().String(), " start reading...")
	defer fmt.Println(con.RemoteAddr().String(), " connect reader exit")
	defer con.Stop()

	for {
		buf := make([]byte, 500)
		_, err := con.Conn.Read(buf)
		if err != nil {
			fmt.Println("Read error ", err)
			con.ExitBufChan <- true
			return
		}

		req := Request{
			connect: con,
			data:    buf,
		}

		go func(request soiface.IRequest) {
			con.Router.PreHandler(request)
			con.Router.Handler(request)
			con.Router.PostHandler(request)
		}(&req)
	}
}

func (con *Connect) Start() {
	go con.StartReader()
	for {
		select {
		case <-con.ExitBufChan:
			return
		}
	}
}

func (con *Connect) Stop() {
	if con.IsClosed == true {
		return
	}

	con.IsClosed = true

	con.Conn.Close()

	con.ExitBufChan <- true

	close(con.ExitBufChan)
}

func (con *Connect) RemoteAddr() net.Addr {
	return con.Conn.RemoteAddr()
}

func (con *Connect) GetConnID() uint32 {
	return con.SessionID
}

func (con *Connect) GetConnection() *net.TCPConn {
	return con.Conn
}
