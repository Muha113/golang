package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	clt, err := net.Dial("tcp", "127.0.0.1:8083")
	if err != nil {
		log.Fatal(err)
	}
	defer clt.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Message to server: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.Compare(message, "exit\n") == 0 {
			break
		}
		fmt.Fprintf(clt, message+"\n")
		receivedMessage, err := bufio.NewReader(clt).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message from server: " + receivedMessage)
	}
}
