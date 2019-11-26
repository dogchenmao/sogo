package sonet

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"strconv"

	"github.com/dogchenmao/sogo/soiface"
)

type DataPack struct{}

//消息封装类实例化
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 加密
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

// []byte 字节切片 循环查找
func SearchByteSliceIndex(bSrc []byte, b byte) int {
	for i := 0; i < len(bSrc); i++ {
		if bSrc[i] == b {
			return i
		}
	}
	return -1
}

// 解密
func AesDecryptECB(encrypted []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(commonkey)
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	bEnd := SearchByteSliceIndex(decrypted, 0)

	return decrypted[:bEnd]
}

//消息封包
func (dp *DataPack) Pack(msg soiface.IMessage) ([]byte, error) {
	var sendData = msg.GetBytes()
	// 头
	var packVer int16 = 3
	var placeHolder int32 = 0
	var dataLen int32 = int32(len(sendData))
	var flag int32 = 3

	// 压缩
	var sendDataEx bytes.Buffer
	w := zlib.NewWriter(&sendDataEx)
	w.Write(msg.GetBytes())
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

	// Crc32
	crc := crc32.ChecksumIEEE(sendData)
	bufData := bytes.NewBuffer(sendData[0:2])
	binary.Write(bufData, binary.LittleEndian, crc)
	bufData.Write(sendData[6:])

	return bufData.Bytes(), nil
}

func (dp *DataPack) UnPack(dataSrc []byte) (soiface.IMessage, error) {
	data := make([]byte, len(dataSrc))
	copy(data, dataSrc)
	// 头
	packVer := binary.LittleEndian.Uint16(data[0:2])
	placeHolder := binary.LittleEndian.Uint32(data[2:6])
	dataLen := binary.LittleEndian.Uint32(data[6:10])
	flag := binary.LittleEndian.Uint32(data[10:14])

	data = data[14:]
	if flag&1 == 1 {
		data = AesDecryptECB(data)
	}
	if flag&2 == 2 {
		b := bytes.NewReader(data)
		var out bytes.Buffer
		r, _ := zlib.NewReader(b)
		io.Copy(&out, r)
		data = out.Bytes()
	}

	if len(data) < 16 {
		return nil, errors.New("len:" + strconv.Itoa(len(data)))
	}

	msg := &Message{
		Data: make([]byte, len(data)-16),
	}
	msg.MsgType = binary.LittleEndian.Uint32(data[0:4])
	msg.RequestID = binary.LittleEndian.Uint32(data[4:8])
	msg.SesID = binary.LittleEndian.Uint32(data[8:12])
	msg.Echo = binary.LittleEndian.Uint32(data[12:16])
	copy(msg.Data, data[16:])

	fmt.Println(packVer, placeHolder, dataLen)
	return msg, nil
}
