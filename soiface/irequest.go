package soiface

type IRequest interface {
	GetConn() IConnect
	GetMsg() IMessage
}
