package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func handleMessage(message string) string {
	if resultStr, err := strconv.ParseInt(message, 10, 64); err == nil {
		resultStr *= int64(2)
		return strconv.Itoa(int(resultStr))
	}
	return strings.ToUpper(message)
}

func main() {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	serv, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		message, err := bufio.NewReader(serv).ReadString('\n')
		if err != nil {
			log.Fatal("Error in bufio.NewReader(serv).ReadString(): ", err)
		}
		fmt.Print("Message received:", string(message))
		splitMessage := strings.Split(message, "\n")
		_, err = serv.Write([]byte(handleMessage(splitMessage[0]) + "\n"))
		if err != nil {
			log.Fatal("Error in serv.Write([]byte(handleMessage(message))): ", err)
		}
	}
}
