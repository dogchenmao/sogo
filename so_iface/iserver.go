package so_iface

type IServer interface {
	Start()
	Stop()
	Serve()
}
