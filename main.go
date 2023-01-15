package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Printf("Hi your local network addr: %s\n", getLocalAddr())
	fmt.Println("Want to connect? y/n")

	input := bufio.NewScanner(os.Stdin)
	connectAddr := ""

	for input.Scan() {
		if input.Text() == "y" {
			fmt.Println("Enter net address to connect")
			connectAddr = getConnectAddr(input)
			break
		}
	}

	if connectAddr != "" {
		go connect(connectAddr)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)

	}
}

func connect(connectAddr string) {
	listener, err := net.Dial("tcp", connectAddr+":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Printf("Connected to %s\n", connectAddr)

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("You: ")
	for input.Scan() {
		fmt.Print("You: ")
		message := fmt.Sprintf(input.Text() + "\n")
		sendMessage(listener, message)
	}
}

func getConnectAddr(scan *bufio.Scanner) string {

	connectAddr := ""
	for scan.Scan() {
		if connectAddr = scan.Text(); connectAddr != "" || net.ParseIP(connectAddr) != nil {
			fmt.Println("Get connect address")
			return connectAddr
		} else {
			fmt.Println("Enter correct ipv4 address")
			getConnectAddr(scan)
		}
	}
	return connectAddr
}

func handleConn(c net.Conn) {
	defer c.Close()
	remoteAddr := c.RemoteAddr()

	fmt.Printf("Client connected: %s\n", remoteAddr)

	for {
		input := make([]byte, 4096)

		_, err := c.Read(input)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(input))
	}
}

func sendMessage(c net.Conn, message string) {
	_, err := c.Write([]byte(message))

	if err != nil {
		log.Fatal(err)
	}

}
