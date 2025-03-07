package main

import (
	"bufio"
	"fmt"
	"net"
)

var (
	server string
	port   string
)

// TODO тестовый клиент
func main() {

	fmt.Print("Enter server ip: ")
	res, err := fmt.Scanf("%s", &server)

	if res != 1 {
		fmt.Printf("Can't read server ip\n")
	}
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter server port: ")
	fmt.Scanf("%s", &port)

	if res != 1 {
		fmt.Printf("Can't read port ip\n")
		return
	}
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", server+":"+port)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(conn)
	for {
		for scanner.Scan() {
			if scanner.Err() != nil {
				conn.Close()
			} else {
				fmt.Println(scanner.Text())
			}
		}
	}

}
