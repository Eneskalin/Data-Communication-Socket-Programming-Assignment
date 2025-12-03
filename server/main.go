package server

import(
	"net"
	"fmt"
)


func main(){
	listener, _ := net.Listen("tcp", ":8000")
	fmt.Println("Server 8000 portundan dinliyor")
	

}