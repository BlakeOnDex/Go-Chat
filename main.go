package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
	
	go Readmsg(conn)
	Sendmsg(conn)
}

func Readmsg(conn net.Conn) {
	for {
		msg := make([]byte,4096)
		Size, err := conn.Read(msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(Size)
		txt := string(msg[:Size])
		fmt.Println("The other person said:", txt)
	}
	conn.Close()
}

func Sendmsg(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		txt := scanner.Bytes()
		_, err := conn.Write(txt)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("You said:",string(txt))

	}
}
