package main 


import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

func main(){

	fmt.Println("Launching the server")
	
	//Listen all participants

	ln, _ := net.Listen("tcp",":8081")

	conn, _ := ln.Accept()

	for true {

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message Received:",message)
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
