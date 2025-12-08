package main


import(
	"fmt"
	"net"


)


func main(){
	var text string;
	fmt.Println("Enter a text:");
	fmt.Scanf("%s",&text);
	connection,err:=net.Dial("tcp","localhost:8000");
	if err !=nil{
		fmt.Println("Server'a baglanti basarisiz oldu",err);
	}
	_,err=connection.Write([]byte(text))
	if(err!=nil){
		fmt.Println("Sunucuya mesaj gonderilemedi")
		return
	}

	connection.Close();
}
