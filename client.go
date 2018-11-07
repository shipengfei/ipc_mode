package ipc_mode

import (
	"encoding/json"
	"github.com/wonderivan/logger"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	return &IpcClient{conn: server.Connect()}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{Method: method, Params: params}
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	client.conn <- string(b)
	res := <-client.conn // 等待服务器返回数据值
	err = json.Unmarshal([]byte(res), resp)
	logger.Info(res)
	return
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
