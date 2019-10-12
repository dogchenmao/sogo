package soiface

type IRequest interface {
	GetConnection() IConnect
	GetData() []byte
}
