package main

import "net"

func main() {
	//监听窗口
	listen, _ := net.Listen("tcp", `:88`)
	for {
		//等待链接连接并建立链接
		conn, _ := listen.Accept()
		//通过链接发送数据
		conn.Write([]byte("我是服务端"))
	}
}
