package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

const BUFFERSIZE = 1024

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

func sendFileToClient(connection net.Conn) {
	fmt.Println("A client has connected!")
	// defer connection.Close()
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	fmt.Println("Sending filename and filesize!")
	print(connection)
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	sendBuffer := make([]byte, BUFFERSIZE)
	fmt.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	fmt.Println("File has been sent, closing connection!")
	return
}

func main() {
	server, err := net.ResolveUDPAddr("udp4", "localhost:27001")
	if err != nil {
		fmt.Println("Error listetning: ", err)
		os.Exit(1)
	}
	// defer server.Close()
	fmt.Println("Server started! Waiting for connections...")
	// for {
	connection, err := net.ListenUDP("udp4", server)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	print(connection)
	fmt.Println("Client connected")
	sendFileToClient(connection)
	// }
}
