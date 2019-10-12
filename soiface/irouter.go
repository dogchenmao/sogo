package soiface

type IRouter interface {
	PreHandler(req IRequest)
	Handler(req IRequest)
	PostHandler(req IRequest)
}
