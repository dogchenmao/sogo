package sonet

import (
	"bytes"
	"encoding/binary"
)

type Message struct {
	MsgType   uint32
	RequestID uint32
	SesID     uint32
	Echo      uint32
	Data      []byte
}

func NewMessage(msgType, requsetId, sesId, echo uint32, data []byte) *Message {
	return &Message{msgType, requsetId, sesId, echo, data}
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) GetBytes() []byte {
	var buf bytes.Buffer

	binary.Write(&buf, binary.LittleEndian, m.MsgType)
	binary.Write(&buf, binary.LittleEndian, m.RequestID)
	binary.Write(&buf, binary.LittleEndian, m.SesID)
	binary.Write(&buf, binary.LittleEndian, m.Echo)
	buf.Write(m.Data)

	return buf.Bytes()
}

func (m *Message) GetMsgId() uint32 {
	return m.RequestID
}
