package chat

import (
	"fmt"
	"log"
	"net"

	webview "github.com/webview/webview_go"
)

func server(w webview.WebView, porta string) {
	server, err := net.Listen("tcp", porta)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go read(w, conn)
	}
}

func read(w webview.WebView, conn net.Conn) {
	for {
		var buffer = make([]byte, 256)
		nbytes, err := conn.Read(buffer)

		if err != nil {
			log.Fatal(err)
		}

		if nbytes > 0 {
			// Nesta parte devo colocar a mensagem com js
			// fmt.Println(conn.RemoteAddr().String() + ">=< : " + string(buffer))
			w.Eval(fmt.Sprintf("postUserMessage('%s')", string(buffer)))
		}

	}
}

func client(w webview.WebView, ip string, porta string) {
	conn, err := net.Dial("tcp", ip+porta)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	w.Bind("send", func(msg string) string {
		_, err = conn.Write([]byte(msg))

		if err != nil {
			log.Fatal(err)
		}

		w.Eval(fmt.Sprintf("postMyMessage('%s')", msg))

		return msg
	})

}
