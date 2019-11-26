package netmgr

import (
	"github.com/dogchenmao/sogo/soiface"
	"github.com/dogchenmao/sogo/sonet"
)

type INetMgr interface {
	SendHallRequest(msgType, requestID, echo uint32, data []byte)
	SendGameRequest(msgType, requestID, echo uint32, data []byte)
}

type NetMgr struct {
	HallClient soiface.IClient
	GameClient soiface.IClient
}

func NewNetMgr(hallHandler, gameHandler soiface.IMsgHandle) *NetMgr {
	//初始化client
	mgr := &NetMgr{}

	mgr.HallClient = sonet.NewClient("192.168.103.108", 31626, hallHandler)
	mgr.HallClient.Start()
	mgr.GameClient = sonet.NewClient("192.168.44.130", 42002, gameHandler)
	mgr.GameClient.Start()

	return mgr
}

func (n *NetMgr) SendHallRequest(msgType, requestID, echo uint32, data []byte) {
	n.HallClient.SendRequest(msgType, requestID, echo, data)
}

func (n *NetMgr) SendGameRequest(msgType, requestID, echo uint32, data []byte) {
	n.GameClient.SendRequest(msgType, requestID, echo, data)
}
