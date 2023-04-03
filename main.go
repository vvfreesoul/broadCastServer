package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	go ListenerForBroadcastResponses()
	go listenerOfGetBroadCastIP()
	//time.Sleep(time.Second * 10000)
}
func ListenerForBroadcastResponses() {
	// 定义本地地址
	fmt.Printf("执行了ListenerForBroadcastResponses")
	localAddr := net.TCPAddr{
		IP:   net.IPv4zero,
		Port: 12345,
	}

	// 创建本地TCP监听器
	listener, err := net.ListenTCP("tcp", &localAddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for broadcast responses on port 12345...")

	for {
		// 接收连接请求
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}

		// 读取数据
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}

		// 显示接收到的数据
		fmt.Printf("Received %d bytes: %s\n", n, string(buffer[:n]))

		// 关闭连接
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}
}

func SendLocalIPToRemoteMachineViaBroadcastIP(remoteIPAddress string) {
	// 获取本机hostname
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// 拼接待发送的消息
	msg := fmt.Sprintf("Hostname: %s", hostname)

	// 解析目标地址
	remoteAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:12345", remoteIPAddress))
	if err != nil {
		panic(err)
	}

	// 连接目标地址
	conn, err := net.DialTCP("tcp", nil, remoteAddr)
	if err != nil {
		panic(err)
	}

	// 发送消息
	_, err = conn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

	// 关闭连接
	err = conn.Close()
	if err != nil {
		panic(err)
	}
}
func listenerOfGetBroadCastIP() {
	fmt.Printf("listenerOfGetBroadCastIP")
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
		SendLocalIPToRemoteMachineViaBroadcastIP(remoteAddr.IP.String())

		// 处理完数据后可以继续等待下一次读取
	}
}
