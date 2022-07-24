# Message Tcp

Set payload length ( Variable-Length Integer Encoding ) in pseudo-header to collect full message

#### Variable-Length Integer Encoding ( rfc9000 )

https://github.com/adrian-lin-1-0-0/go-variable-length-integer-encoding

#### Checksum (Optional)

- 1's Complement


## Install

```
go get github.com/adrian-lin-1-0-0/go-message-tcp@latest

```

## Usage

Client

```go
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
```

Server

```go
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
```
