package soiface

type IMsgHandle interface {
	DoMsgHandler(request IRequest)
	AddRouter(router IRouter)
	StartWorkPool()
	SendMsgToTaskQueue(request IRequest)
}
