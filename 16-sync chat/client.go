package main

import (
    "fmt"
	"net"
	"os"
	"bufio"
	"log"
)

const (
    connection_host = "localhost"
    connection_port = "8080"
    connection_type = "tcp"
)

func main() {
    fmt.Println("Connecting to " + connection_type + " server " + connection_host + ":" + connection_port)
	
	connection, error := net.Dial(connection_type, connection_host + ":" + connection_port)
   
	if error != nil {
        fmt.Println("Error connecting:", error.Error())
        os.Exit(1)
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Text to send: ")

        input, _ := reader.ReadString('\n')

        connection.Write([]byte(input))

        message, _  := bufio.NewReader(connection).ReadString('\n')

        log.Print("Server relay:", message)
    }
}
