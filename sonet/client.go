package sonet

import (
	"github.com/dogchenmao/sogo/game_my"
	"github.com/golang/protobuf/proto"
)

func MyFunc() error {
	f := func(num int) *int32 {
		q := int32(num)
		return &q
	}

	f2 := func(num string) *string {
		q := num
		return &q
	}

	req := &game_my.EnterNormalGameReq{
		Userid:      f(13061),
		Gameid:      f(1002),
		Roomid:      f(7866),
		Flag:        f(0),
		Target:      f(0),
		Hardid:      f2(""),
		HallVersion: f2(""),
		GameVersion: f2(""),
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return nil
	}
	prase_data := &game_my.EnterNormalGameReq{}

	err = proto.Unmarshal(data, prase_data)

	return nil
}
