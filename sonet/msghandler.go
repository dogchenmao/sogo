package sonet

import (
	"github.com/dogchenmao/sogo/soiface"
)

type MsgHandler struct {
	Routers        soiface.IRouter
	WorkerPoolSize uint32
	TaskQueue      []chan soiface.IRequest
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		WorkerPoolSize: 32,
		TaskQueue:      make([]chan soiface.IRequest, 32),
	}
}

func (mh *MsgHandler) AddRouter(router soiface.IRouter) {
	// if _, ok := mh.Routers[msgId]; ok {
	// 	panic("repeated msg handler, msgid = " + strconv.Itoa(int(msgId)))
	// }

	mh.Routers = router
}

func (mh *MsgHandler) DoMsgHandler(req soiface.IRequest) {
	// handler, ok := mh.Routers[req.GetMsg().GetMsgId()]
	// if !ok {
	// 	return
	// }

	mh.Routers.Handler(req)
}

func (mh *MsgHandler) SendMsgToTaskQueue(req soiface.IRequest) {

}

func (mh *MsgHandler) StartOneWorker(workerId int, taskQueue chan soiface.IRequest) {
	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandler) StartWorkPool() {
	for index := 0; index < int(mh.WorkerPoolSize); index++ {
		mh.TaskQueue[index] = make(chan soiface.IRequest, 1024)

		go mh.StartOneWorker(index, mh.TaskQueue[index])
	}
}
