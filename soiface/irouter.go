package soiface

type IRouter interface {
	Handler(req IRequest)
}
