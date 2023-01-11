package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	address, _ := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	listener, _ := net.ListenMulticastUDP("udp", nil, address)
	defer listener.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddress, _ := listener.ReadFromUDP(buffer)
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now    %s\n", string(buffer[:length]))
	}
}
