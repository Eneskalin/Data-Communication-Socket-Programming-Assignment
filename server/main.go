package main

import(
	"net"
	"fmt"
	"server/handlers"
)


func main(){
	listener, _ := net.Listen("tcp", ":8000")
	fmt.Println("Server 8000 portundan dinliyor")
	
	client2Conn, err := net.Dial("tcp", "localhost:9000")
	if err!=nil {
		fmt.Println("Receiver 9000 portunda bulunamadı");
		return;
	}
	defer client2Conn.Close()
	for {
		conn, _ := listener.Accept()
		go handlers.HandleConnection(conn, client2Conn)
	}



}