package ipc_mode

import (
	"encoding/json"
	// "github.com/wonderivan/logger"
)

type Server interface {
	Name() string
	Handle(method string, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{Server: server}
}

// 返回一个channel，用户连接成功后，发送数据请求
func (i_server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(ch chan string) {
		for {
			message := <-ch
			if message == "CLOSE" {
				break
			}

			var request Request
			err := json.Unmarshal([]byte(message), &request)
			if err != nil {

			}

			resp := i_server.Handle(request.Method, request.Params)

			b, err := json.Marshal(resp)
			ch <- string(b)
		}
	}(session)

	return session
}
