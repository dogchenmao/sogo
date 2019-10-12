package soiface

type IServer interface {
	Start()
	Stop()
	Serve()
}
