package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器 通信协议,ip地址,port
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待客户端连接...")
	//阻塞监听客户端请求
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("lintener.Accpet() err:", err)
		return
	}
	defer conn.Close()

	fmt.Println("连接成功!")

	//读取数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	//处理数据
	fmt.Println("收到数据:", string(buf[:n]))

}
