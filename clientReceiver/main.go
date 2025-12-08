package main

import(
	"net"
	"fmt"
)


func main(){
	ln,err:=net.Listen("tcp",":9000");
	if err !=nil{
		fmt.Println("Sunucu 9000 portunda baslatilamadi");
		return
	}else{
		fmt.Println("Sunucu 9000 portunda dinliyor")
	}
	for{
		conn,err:=ln.Accept();
		if err!=nil{
			fmt.Println(err)
		}
		go HandleConnection(conn)

	}

}

func HandleConnection(conn net.Conn) {
    
    defer conn.Close()

    
    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println(err)
        return
    }

    
    fmt.Printf("Received: %s", buf)
}