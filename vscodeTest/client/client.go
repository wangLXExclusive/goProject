package main

import (
	"fmt"
	"net"
)

func main() {
	//请求链接
	conn, _ := net.Dial("tcp", "192.168.5.6:88")
	//数据容器
	buf := make([]byte, 2048)
	//读取数据到容器
	l, _ := conn.Read(buf)
	//打印输出
	fmt.Println(string(buf[:l]))
}
