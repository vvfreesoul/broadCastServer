package main

import (
	"fmt"
	"net"
)

func main() {
	// 定义本地地址
	localAddr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 12345,
	}

	// 创建本地UDP连接
	conn, err := net.ListenUDP("udp", &localAddr)
	if err != nil {
		panic(err)
	}

	// 读取数据
	buffer := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		panic(err)
	}

	// 显示接收到的数据和发送方地址
	fmt.Printf("Received %d bytes from %v: %s\n", n, remoteAddr, string(buffer[:n]))

	// 关闭连接
	err = conn.Close()
	if err != nil {
		panic(err)
	}
}
