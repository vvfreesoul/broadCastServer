package main

import "net"

func main() {
	broadcastMyIP()
}

func broadcastMyIP() {
	// 定义广播地址
	broadcastAddr := net.IPv4(255, 255, 255, 0)

	// 定义本地地址
	localAddr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 0,
	}

	// 创建本地UDP连接
	conn, err := net.DialUDP("udp", &localAddr, &net.UDPAddr{
		IP:   broadcastAddr,
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}

	// 发送数据到广播地址
	_, err = conn.Write([]byte("Hello, world!"))
	if err != nil {
		panic(err)
	}

	// 关闭连接
	err = conn.Close()
	if err != nil {
		panic(err)
	}
}
