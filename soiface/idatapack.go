package soiface

type IDataPack interface {
	Pack(msg IMessage) ([]byte, error)
	UnPack([]byte) (IMessage, error)
}
