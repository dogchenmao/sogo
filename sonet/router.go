package sonet

import (
	"github.com/dogchenmao/sogo/soiface"
)

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(req soiface.IRequest)  {}
func (br *BaseRouter) Handle(req soiface.IRequest)     {}
func (br *BaseRouter) PostHandle(req soiface.IRequest) {}
