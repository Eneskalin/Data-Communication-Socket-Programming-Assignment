package main

import (
	"fmt"
	"net"
	"server/handlers"
	"server/helpers"
	"time"
)

func main() {
	config := helpers.LoadConfig()
	listener, _ := net.Listen("tcp", ":"+config.Ports.SenderPort)
	fmt.Println("[Server]", config.Ports.SenderPort, "portundan dinliyor")

	var client2Conn net.Conn
	var err error

	
	go func() {
		for {
			fmt.Println("[Server] Receiver", config.Ports.ReceiverPort, "portunda bağlanılıyor.")
			client2Conn, err = net.Dial("tcp", "localhost:"+config.Ports.ReceiverPort)
			if err != nil {
				fmt.Println("[Server] Receiver,bağlantısı başarısız.")
				time.Sleep(5 * time.Second)
			} else {
				
				break
			}
		}
	}()

	for {
		conn, _ := listener.Accept()
		if client2Conn != nil {
			go handlers.HandleConnection(conn, client2Conn)
		}
	}
}
