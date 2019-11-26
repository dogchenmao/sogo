package robot

import (
	"github.com/dogchenmao/sogo/pb/game_my"
	"github.com/dogchenmao/sogo/soiface"
	"github.com/dogchenmao/sogo/sonet"
	"github.com/golang/protobuf/proto"
)

type GameRouter struct {
	sonet.BaseRouter
}

func (g *GameRouter) Handler(req soiface.IRequest) {
	switch req.GetMsg().GetMsgId() {
	case 10020001:
		resp := &game_my.EnterNormalGameResp{}
		proto.Unmarshal(req.GetMsg().GetData(), resp)
		break
	case 10020033:
		break
	}
}
