package sonet

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"hash/crc32"

	"github.com/dogchenmao/sogo/soiface"
)

type DataPack struct{}

//消息封装类实例化
func NewDataPack() *DataPack {
	return &DataPack{}
}

//消息封包
func (dp *DataPack) Pack(msg soiface.IMessage) ([]byte, error) {
	var sendData = msg.GetData()
	// 头
	var packVer int16 = 3
	var placeHolder int32 = 0
	var dataLen int32 = int32(len(sendData))
	var flag int32 = 3

	// 压缩
	var sendDataEx bytes.Buffer
	w := zlib.NewWriter(&sendDataEx)
	w.Write(msg.GetData())
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

func (dp *DataPack) UnPack([]byte) (soiface.IMessage, error) {
	return nil, nil
}
