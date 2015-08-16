package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func main(){
	
	conn,_ := net.Dial("tcp","127.0.0.1:8081")
	for true {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to Send:")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")
		message,_ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}



