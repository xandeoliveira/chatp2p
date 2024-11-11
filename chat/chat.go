package chat

import (
	"fmt"
	"log"
	"net"
)

func server(porta string) {
	server, err := net.Listen("tcp", porta)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go read(conn)
	}
}

func read(conn net.Conn) {
	for {
		var buffer = make([]byte, 256)
		nbytes, err := conn.Read(buffer)

		if err != nil {
			log.Fatal(err)
		}

		if nbytes > 0 {
			// Nesta parte devo colocar a mensagem com js
			// fmt.Println(conn.RemoteAddr().String() + ">=< : " + string(buffer))
			fmt.Printf("postUserMessage('%s')", string(buffer))
		}

	}
}

func client(ip string) net.Conn {
	conn, err := net.Dial("tcp", ip+":9000")

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
