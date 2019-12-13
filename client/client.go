package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	clt, _ := net.Dial("tcp", "127.0.0.1:8083")
	defer clt.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Message to server: ")
		message, _ := reader.ReadString('\n')
		if strings.Compare(message, "exit\n") == 0 {
			break
		}
		fmt.Fprintf(clt, message+"\n")
		receivedMessage, _ := bufio.NewReader(clt).ReadString('\n')
		fmt.Print("Message from server: " + receivedMessage)
	}
}
