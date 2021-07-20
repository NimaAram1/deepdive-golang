package main 

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"log"	
)

const (
	connection_host = "localhost"
	connection_port = "8080"
	connection_type = "tcp" 
)

func main(){
	fmt.Println("Starting " + connection_type + "server on " + connection_host + ":" + connection_port)
	server, error := net.Listen(connection_type, connection_host + ":" + connection_port)
	
	if error != nil {
		fmt.Println("Error listening: ", error.Error())
		os.Exit(1)
	}
	
	defer server.Close()
	
	for {
		client, error := server.Accept()
		if error != nil {
			fmt.Println("Error accepting: ", error.Error())
			return
		}
		fmt.Println("Client connected.")
		fmt.Println("Client " + client.RemoteAddr().String() + "connected.")

		go handleConnection(client)
	}
}

func handleConnection(connection net.Conn){
	buffer, err := bufio.NewReader(connection).ReadBytes('\n')

    if err != nil {
        fmt.Println("Client left.")
        connection.Close()
        return
    }

    log.Println("Client message:", string(buffer[:len(buffer)-1]))

    connection.Write(buffer)

    handleConnection(connection)
}