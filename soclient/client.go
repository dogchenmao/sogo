package soclient

// import (
// 	"bytes"
// 	"compress/zlib"
// 	"crypto/aes"
// 	"encoding/binary"
// 	"flag"
// 	"hash/crc32"
// 	"log"
// 	"net/url"
// 	"time"

// 	"github.com/dogchenmao/sogo/soiface"
// 	"github.com/gorilla/websocket"
// )

// type Client struct {
// 	connect soiface.IConnect
// }

// var (
// 	commonkey = []byte{6, 14, 1, 8, 16, 8, 3, 2, 13, 2, 20, 15, 19, 4, 7, 10}
// )

// func AesEncryptECB(origData []byte) (encrypted []byte) {
// 	cipher, _ := aes.NewCipher(commonkey)
// 	length := (len(origData) + aes.BlockSize) / aes.BlockSize
// 	plain := make([]byte, length*aes.BlockSize)
// 	copy(plain, origData)

// 	encrypted = make([]byte, len(plain))
// 	// 分组分块加密
// 	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
// 		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
// 	}

// 	return encrypted
// }

// var addr = flag.String("addr", "192.168.103.153:42002", "http service address")

// func NewClient(host string, port int) {

// 	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
// 	log.Printf("connecting to %s  %d", u.String(), port)
// 	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		log.Fatal("dial:", err)
// 	}
// 	defer ws.Close()
// 	done := make(chan struct{})
// 	go func() {
// 		defer close(done)
// 		for {
// 			_, message, err := ws.ReadMessage()
// 			if err != nil {
// 				log.Println("read:", err)
// 				return
// 			}
// 			log.Printf("recv: %s", message)
// 		}
// 	}()

// 	resp := NewResp(0, 10020033, 1, 1)
// 	sendData := resp.GetBytes()

// 	// 头
// 	var packVer int16 = 3
// 	var placeHolder int32 = 0
// 	var dataLen int32 = int32(len(sendData))
// 	var flag int32 = 3

// 	// 压缩
// 	var sendDataEx bytes.Buffer
// 	w := zlib.NewWriter(&sendDataEx)
// 	w.Write(sendData)
// 	w.Close()

// 	// 加密
// 	sendData = AesEncryptECB(sendDataEx.Bytes())

// 	sendDataEx.Reset()
// 	binary.Write(&sendDataEx, binary.LittleEndian, packVer)
// 	binary.Write(&sendDataEx, binary.LittleEndian, placeHolder)
// 	binary.Write(&sendDataEx, binary.LittleEndian, dataLen)
// 	binary.Write(&sendDataEx, binary.LittleEndian, flag)

// 	// 拼接数据
// 	sendData = append(sendDataEx.Bytes(), sendData...)

// 	// Crc32
// 	crc := crc32.ChecksumIEEE(sendData)
// 	bufData := bytes.NewBuffer(sendData[0:2])
// 	binary.Write(bufData, binary.LittleEndian, crc)
// 	bufData.Write(sendData[6:])

// 	for {
// 		err := ws.WriteMessage(websocket.BinaryMessage, bufData.Bytes())
// 		if err != nil {
// 			log.Println("write:", err)
// 			return
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }
