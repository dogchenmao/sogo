package sonet

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/dogchenmao/sogo/soiface"

	"github.com/dogchenmao/sogo/pb/game_my"
	"github.com/golang/protobuf/proto"

	"github.com/gorilla/websocket"
)

var (
	commonkey = []byte{6, 14, 1, 8, 16, 8, 3, 2, 13, 2, 20, 15, 19, 4, 7, 10}
)

type Connect struct {
	Conn        *websocket.Conn
	SessionID   uint32
	IsClosed    bool
	msgHandler  soiface.IMsgHandle
	msgChan     chan []byte
	ExitBufChan chan bool
}

func NewConnect(ip string, port int, handler soiface.IMsgHandle) *Connect {
	addr := ip + ":" + strconv.Itoa(port)
	u := url.URL{Scheme: "ws", Host: addr, Path: "/echo"}
	log.Printf("connecting to %s  %d", u.String(), port)
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	s := &Connect{
		Conn:        ws,
		SessionID:   1,
		IsClosed:    true,
		msgHandler:  handler,
		msgChan:     make(chan []byte),
		ExitBufChan: make(chan bool, 1),
	}
	return s
}

// 启动连接开始工作
func (con *Connect) Start() {
	go con.StartReader()
	go con.StartWrite()
	// con.SendPulse()
}

func (con *Connect) StartWrite() {
	for {
		select {
		case data := <-con.msgChan:
			err := con.Conn.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}

func (con *Connect) StartReader() {
	defer con.Stop()
	for {
		dp := NewDataPack()
		_, data, err := con.Conn.ReadMessage()
		if err != nil {
			return
		}
		msg, err := dp.UnPack(data)
		if err != nil {
			return
		}

		req := &Request{
			connect: con,
			msg:     msg,
		}
		//客户端直接处理
		go con.msgHandler.DoMsgHandler(req)
	}
}

func (con *Connect) Stop() {

}

func (con *Connect) SendNotify() {

}

func (con *Connect) SendRequest(MsgType, RequestID, Echo uint32, data []byte) {
	con.SessionID++
	if Echo == 1 {
		msg := NewMessage(MsgType, RequestID, con.SessionID, Echo, data)
		data, _ := NewDataPack().Pack(msg)
		con.msgChan <- data
	} else {

	}
}

// func (con *Connect) SendData(sendData []byte) {
// 	err := con.Conn.WriteMessage(websocket.BinaryMessage, sendData)
// 	if err != nil {
// 		log.Println("write:", err)
// 		return
// 	}
// }

func (con *Connect) SendPulse() {
	f := func(num int) *int32 {
		q := int32(num)
		return &q
	}

	f2 := func(num string) *string {
		q := num
		return &q
	}

	for {
		req := &game_my.EnterNormalGameReq{
			Userid:      f(13061),
			Gameid:      f(1002),
			Roomid:      f(7866),
			Flag:        f(0),
			Target:      f(0),
			Hardid:      f2(""),
			HallVersion: f2(""),
			GameVersion: f2(""),
		}

		data, err := proto.Marshal(req)
		if err != nil {
			return
		}

		con.SendRequest(0, 10020001, 1, data)
		time.Sleep(1 * time.Second)
	}
}

func (con *Connect) Close() {}

func (con *Connect) AddRouter(router soiface.IRouter) {
	con.msgHandler.AddRouter(router)
}
