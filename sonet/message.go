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
	// Data      []byte
}

func NewMessage(msgType, requsetId, sesId, echo uint32) *Message {
	return &Message{msgType, requsetId, sesId, echo}
}

func (m *Message) GetData() []byte {
	var buf bytes.Buffer

	binary.Write(&buf, binary.LittleEndian, m.MsgType)
	binary.Write(&buf, binary.LittleEndian, m.RequestID)
	binary.Write(&buf, binary.LittleEndian, m.SesID)
	binary.Write(&buf, binary.LittleEndian, m.Echo)

	return buf.Bytes()
}
