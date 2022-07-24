package main

import (
	"log"

	tcp "github.com/adrian-lin-1-0-0/go-message-tcp"
)

func main() {
	address := ":5000"
	conn, err := tcp.Dial(address)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	conn.SetBufLen(2048)   //default 1024
	conn.SetChecksum(true) //default false

	conn.Write([]byte("hello server"))

	buf, err := conn.Read()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("receive data : ", string(buf))

}
