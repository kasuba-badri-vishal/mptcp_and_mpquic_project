package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"reflect"
	"time"
)

const BUFFERSIZE = 1024

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

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

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	print(reflect.TypeOf(connection))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	// buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	// for {
	// 	fmt.Println("A client has connected!")
	// 	file, err := os.Open("example.txt")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fileInfo, err := file.Stat()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	// 	fileName := fillString(fileInfo.Name(), 64)
	// 	fmt.Print(fileSize, fileName)
	// 	fmt.Println("Sending filename and filesize!")
	// 	connection.Write([]byte(fileSize))
	// 	connection.Write([]byte(fileName))
	// 	sendBuffer := make([]byte, BUFFERSIZE)
	// 	fmt.Println("Start sending file!")
	// 	for {
	// 		_, err = file.Read(sendBuffer)
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		connection.Write(sendBuffer)
	// 	}
	// 	fmt.Println("File has been sent, closing connection!")
	// 	n, addr, err := connection.ReadFromUDP(buffer)
	// 	fmt.Print("-> ", string(buffer[0:n-1]))
	// 	if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
	// 		fmt.Println("Exiting UDP server!")
	// 		return
	// 	}

	// 	data := []byte(strconv.Itoa(random(1, 1001)))
	// 	fmt.Printf("data: %s\n", string(data))
	// 	_, err = connection.WriteToUDP(data, addr)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }
}
