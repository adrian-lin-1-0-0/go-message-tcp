package main

import (
	"log"

	tcp "github.com/adrian-lin-1-0-0/go-message-tcp"
)

func main() {
	address := ":5000"
	li, err := tcp.Listen(address)
	log.Println("server listen on ", address)
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	li.SetBufLen(2048)   //default 1024
	li.SetChecksum(true) // default false

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn tcp.Conn) {
	for {
		buf, err := conn.Read()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("receive data : ", string(buf))
		_, err = conn.Write([]byte("hi client"))
		if err != nil {
			log.Println(err)
			break
		}
	}

	defer conn.Close()

	log.Println("Connection close")
}
