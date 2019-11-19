package sonet

import (
	"bytes"
	"crypto/aes"
	"encoding/binary"
	"flag"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type Resp struct {
	MsgType   int32
	RequestID int32
	SesID     int32
	Echo      int32
}

func NewResp(MsgType, RequestID, SesID, Echo int32) *Resp {
	return &Resp{MsgType, RequestID, SesID, Echo}
}

func (resp *Resp) GetBytes() []byte {
	var buf bytes.Buffer

	binary.Write(&buf, binary.LittleEndian, resp.MsgType)
	binary.Write(&buf, binary.LittleEndian, resp.RequestID)
	binary.Write(&buf, binary.LittleEndian, resp.SesID)
	binary.Write(&buf, binary.LittleEndian, resp.Echo)

	return buf.Bytes()
}

var (
	commonkey = []byte{6, 14, 1, 8, 16, 8, 3, 2, 13, 2, 20, 15, 19, 4, 7, 10}
)

func AesEncryptECB(origData []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(commonkey)
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)

	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

type Connect struct {
	Conn        *websocket.Conn
	SessionID   uint32
	IsClosed    bool
	ExitBufChan chan bool
}

func NewConnect(ip string, port int) *Connect {
	var addr = flag.String("addr", ip+":"+strconv.Itoa(port), "http service address")

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s  %d", u.String(), port)
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	s := &Connect{
		Conn:        ws,
		SessionID:   1,
		IsClosed:    true,
		ExitBufChan: make(chan bool, 1),
	}
	return s
}

func (con *Connect) Start() {
	go con.SendPulse()
}

func (con *Connect) SendNotify() {

}

func (con *Connect) SendRequest(MsgType, RequestID, Echo uint32) {
	con.SessionID++
	if Echo == 1 {
		msg := NewMessage(MsgType, RequestID, con.SessionID, Echo)
		data, _ := NewDataPack().Pack(msg)
		con.SendData(data)
	} else {

	}
}

func (con *Connect) SendData(sendData []byte) {
	err := con.Conn.WriteMessage(websocket.BinaryMessage, sendData)
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func (con *Connect) SendPulse() {
	for {
		con.SendRequest(0, 10020033, 1)
		time.Sleep(1 * time.Second)
	}
}

func (con *Connect) Close() {}

func (con *Connect) HandlerMessage() {}
