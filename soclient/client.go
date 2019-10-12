package soclient

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"encoding/binary"
	"flag"
	"hash/crc32"
	"log"
	"net/url"
	"time"

	"github.com/dogchenmao/sogo/soiface"
	"github.com/gorilla/websocket"
)

type Client struct {
	connect soiface.IConnect
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

var addr = flag.String("addr", "localhost:42002", "http service address")

func NewClient(host string, port int) {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer ws.Close()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	var msgType int32 = 0
	var reqID int32 = 10020033
	var SesID int32 = 0
	var Echo int32 = 1

	var buf bytes.Buffer

	binary.Write(&buf, binary.LittleEndian, msgType)
	binary.Write(&buf, binary.LittleEndian, reqID)
	binary.Write(&buf, binary.LittleEndian, SesID)
	binary.Write(&buf, binary.LittleEndian, Echo)

	// sendBuf := new(bytes.Buffer)
	// binary.Write(sendBuf, binary.LittleEndian, packVer)
	// binary.Write(sendBuf, binary.LittleEndian, param)
	// binary.Write(sendBuf, binary.LittleEndian, dataLen)
	// binary.Write(sendBuf, binary.LittleEndian, flag)

	// // 拼接发送内容
	// binary.Write(sendBuf, binary.LittleEndian, se)

	// sendBuf.Read(msg2)

	// wwwwww := []byte{3, 0, 210, 53, 173, 66, 16, 0, 0, 0, 3, 0, 0, 0, 5, 212, 0, 126, 45, 181, 56, 215, 202, 86, 158, 96, 92, 144, 135, 192, 172, 210, 117, 197, 152, 67, 115, 34, 80, 222, 197, 91, 112, 122, 194, 175}

	sendData := []byte{0, 0, 0, 0, 193, 228, 152, 0, 52, 0, 0, 0, 1, 0, 0, 0}

	// 头
	var packVer int16 = 3
	var placeHolder int32 = 0
	var dataLen int32 = int32(len(sendData))
	var flag int32 = 3

	// 压缩
	var sendDataEx bytes.Buffer
	w := zlib.NewWriter(&sendDataEx)
	w.Write(sendData)
	w.Close()

	// 加密
	sendData = AesEncryptECB(sendDataEx.Bytes())

	sendDataEx.Reset()
	binary.Write(&sendDataEx, binary.LittleEndian, packVer)
	binary.Write(&sendDataEx, binary.LittleEndian, placeHolder)
	binary.Write(&sendDataEx, binary.LittleEndian, dataLen)
	binary.Write(&sendDataEx, binary.LittleEndian, flag)

	// 拼接数据
	sendData = append(sendDataEx.Bytes(), sendData...)

	crc := crc32.ChecksumIEEE(sendData)

	bufData := bytes.NewBuffer(sendData[0:2])
	binary.Write(bufData, binary.LittleEndian, crc)

	bufData.Write(sendData[6:])

	for {
		err := ws.WriteMessage(websocket.BinaryMessage, bufData.Bytes())
		if err != nil {
			log.Println("write:", err)
			return
		}
		time.Sleep(1 * time.Second)
	}

	return
}
