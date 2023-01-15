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
	input := bufio.NewScanner(os.Stdin)

	nickName := ""

	fmt.Println("Enter your nickname")
	for input.Scan() {
		nickName = input.Text()
		break
	}

	fmt.Println("Want to connect? y/n")
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
		connect(conn, nickName)
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

			connect(conn, nickName)
		}
	}
}

func connect(conn net.Conn, nickName string) {

	done := make(chan struct{})

	go func() {
		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			sendMessage(conn, input.Text(), nickName)
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

func sendMessage(c net.Conn, message string, nickName string) {

	fmt.Printf("You: %s\n", message)
	message = fmt.Sprintf("%s: %s\n", nickName, message)
	_, err := c.Write([]byte(message))

	if err != nil {
		log.Fatal(err)
	}

}
