package soiface

type IClient interface {
	Start()
	SendRequest(msgType, requestID, echo uint32, data []byte)
}
