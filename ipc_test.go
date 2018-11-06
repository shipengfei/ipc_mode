package ipc_mode

import (
	// "github.com/shipengfei/ipc_mode"
	"github.com/wonderivan/logger"
	"testing"
)

type EchoServer struct {
}

func (e *EchoServer) Handle(method, params string) string {
	return "ECHO:" + params
}

func (e *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	logger.Info("...")
	server := NewIpcServer(EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	client1.Call("A", "From Client1")
	client2.Call("B", "From Client2")

	client1.Close()
	client2.Close()
}
