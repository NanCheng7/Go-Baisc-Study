package main

import (
	"fmt"
	"net"
)

// UDP客户端
func main() {
	socket, err := net.DialUDP(
		"udp",
		nil,
		&net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: 30000,
		},
	)
	if err != nil {
		fmt.Println("Connection failed,err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("NanCheng Server")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("Send data failed,err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("recieve data failed,err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
