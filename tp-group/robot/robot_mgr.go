package robot

import (
	"time"

	"github.com/dogchenmao/sogo/pb/game_my"
	"github.com/dogchenmao/sogo/pb/hall_base"
	"github.com/dogchenmao/sogo/sonet"
	"github.com/dogchenmao/sogo/tp-group/netmgr"
	"github.com/golang/protobuf/proto"
)

func f(num int) *int32 {
	q := int32(num)
	return &q
}

func f2(num string) *string {
	q := num
	return &q
}

type RobotMgr struct {
	Robots map[int]netmgr.INetMgr
}

func NewRobotMgr() *RobotMgr {
	return &RobotMgr{
		Robots: make(map[int]netmgr.INetMgr, 1024),
	}
}

func (r *RobotMgr) Init() {
	//构造路由
	hallMsgHandler := sonet.NewMsgHandler()
	hallMsgHandler.AddRouter(&HallRouter{})
	gameMsgHandler := sonet.NewMsgHandler()
	gameMsgHandler.AddRouter(&GameRouter{})

	for i := 0; i < 100; i++ {
		r.Robots[i] = netmgr.NewNetMgr(hallMsgHandler, gameMsgHandler)
	}

	go r.test()
}

func (r *RobotMgr) test() {
	r.Login()
	time.Sleep(time.Second * 10)
	r.SendEnterGame()
}

func (r *RobotMgr) SendEnterGame() {
	req := &game_my.EnterNormalGameReq{
		Userid:      f(709854),
		Gameid:      f(1002),
		Roomid:      f(7858),
		Flag:        f(0),
		Target:      f(0),
		Hardid:      f2("50A4C8172ce70000000000000000000"),
		HallVersion: f2("888.8.8"),
		GameVersion: f2("888.8.8"),
	}
	data, err := proto.Marshal(req)
	if err != nil {
		return
	}

	r.Robots[0].SendGameRequest(1, 10020001, 1, data)
}

func (r *RobotMgr) Login() {
	req := &hall_base.LogonReq{
		Userid:        f(709854),
		Password:      f2("1q2w3e"),
		Hardid:        f2("50A4C8172ce70000000000000000000"),
		Systemid:      f2("2891bf8c2c172ce7"),
		Imei:          f2("354096447172ce7"),
		Deviceid:      f2(""),
		Gameid:        f(0),
		Channelid:     f(0),
		Groupid:       f(1),
		Hallsvrid:     f(0),
		Flag:          f(5120),
		LogonInterval: f(1),
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return
	}

	r.Robots[0].SendHallRequest(1, 10010202, 1, data)
}
