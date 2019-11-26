package sonet

import (
	"github.com/dogchenmao/sogo/soiface"
)

type BaseRouter struct{}

func (br *BaseRouter) Handle(req soiface.IMessage) {}
