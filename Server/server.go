package main

import (
	"fmt"
	"github.com/710112257/Zinx/ziface"
	"github.com/710112257/Zinx/znet"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(1, []byte("我是服务器"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("ydc")
	s.AddRouter(&PingRouter{})

	s.Serve()

	select {}
}
