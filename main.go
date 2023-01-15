package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	fmt.Printf("Hi your local network addr: %s\n", getLocalAddr())
	fmt.Println("Want to connect? y/n")

	input := bufio.NewScanner(os.Stdin)
	connectAddr := ""

	for input.Scan() {
		if input.Text() == "y" {
			fmt.Println("Enter net address to connect")
			connectAddr = getConnectAddr(input)
		}
		break
	}

	if connectAddr != "" {
		conn, err := net.Dial("tcp", connectAddr+":8000")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		fmt.Printf("Connected to %s\n", connectAddr)
		connect(conn)
	} else {
		listener, err := net.Listen("tcp", ":8000")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Waiting for connection...")
		defer listener.Close()
		for {

			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}

			connect(conn)
		}
	}
}

func connect(conn net.Conn) {

	done := make(chan struct{})

	go func() {
		input := bufio.NewScanner(os.Stdin)

		fmt.Print("You: ")
		for input.Scan() {
			fmt.Print("You: ")
			message := fmt.Sprintf(input.Text() + "\n")
			sendMessage(conn, message)
		}

		done <- struct{}{}
	}()

	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Fatal(err)
			done <- struct{}{}
		}
		done <- struct{}{}
	}()
	<-done

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

func sendMessage(c net.Conn, message string) {
	_, err := c.Write([]byte(message))

	if err != nil {
		log.Fatal(err)
	}

}
