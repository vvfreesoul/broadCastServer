package main

import (
	"fmt"
	"net"
)

func main() {
	listenerOfGetBroadCastIP()
}
func SendLocalIPToRemoteMachineViaBroadcastIP(remoteAddress string) {

}
func listenerOfGetBroadCastIP() {
	// 定义本地地址
	localAddr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 1234,
	}

	// 创建本地UDP连接
	conn, err := net.ListenUDP("udp", &localAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		// 读取数据
		buffer := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error while reading from UDP: %v\n", err)
			continue
		}

		// 处理接收到的数据
		fmt.Printf("Received %d bytes from %v: %s\n", n, remoteAddr, string(buffer[:n]))

		// 处理完数据后可以继续等待下一次读取
	}
}
